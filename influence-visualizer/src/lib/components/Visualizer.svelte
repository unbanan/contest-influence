<script lang="ts">
	import { resolve } from '$app/paths';
	export let n;
	export let m;
	export let fieldMask: boolean[];
	export let start_positions: Position[];
	export let big_cells: Position[];
	export let rounds: Round[];
	export let names: string[];

	let container: HTMLElement;

	interface Cell {
		row: number;
		col: number;
		value: number;
	}

	interface AttackMove {
		from: Cell;
		to: Cell;
		is_winner: boolean;
		statistics: Statistics;
	}

	interface Attack {
		moves: AttackMove[];
	}

	interface Defence {
		cells: Cell[];
		statistics: Statistics;
	}

	interface Round {
		attack: Attack;
		defence: Defence;
	}

	interface Position {
		row: number;
		col: number;
	}
	const padding = 8;
	const h_n = n;
	const w_n = m;
	const player_number = start_positions.length;
	import { onMount } from 'svelte';

	$: if (grid && ticks[currentTickIndex]) {
		showTick(ticks[currentTickIndex], grid, two);
		if (typeof two !== 'undefined') two.update();
	}

	export let currentScores: Record<string, number> = {};
	$: if (ticks[currentTickIndex]) {
		currentScores = ticks[currentTickIndex].statistics.scores;
	}

	const colors: number[] = [];
	for (let i = 0; i < player_number; i++) {
		colors.push((360 / (player_number + 2)) * (i + 1));
	}

	import Two from 'two.js';
	import type { Group } from 'two.js/src/group';
	import type { Polygon } from 'two.js/src/shapes/polygon';
	import type { Text as TwoText } from 'two.js/src/text';

	interface GridCell {
		group: Group;
		hex: Polygon;
		inner: Polygon;
		text: TwoText;
	}

	interface Position {
		row: number;
		col: number;
	}

	interface FieldCell {
		color: number;
		position: Position;
		value: number;
	}

	interface Statistics {
		scores: Record<string, number>;
	}

	interface Field {
		field: FieldCell[][];
		statistics: Statistics;
	}

	const ticks: Field[] = [];
	let two: Two;
	let params;
	let R, r, sides, text_size;
	let field;
	let grid: GridCell[][];

	export let currentTickIndex = 0;
	let isPlaying = false;

	$: if (grid && ticks[currentTickIndex]) {
		showTick(ticks[currentTickIndex], grid, two);
		if (typeof two !== 'undefined') two.update();
	}

	onMount(() => {
		params = {
			autostart: true,
			height: container.clientHeight,
			width: container.clientWidth,
			type: Two.Types.canvas
		};
		two = new Two(params).appendTo(container);

		// R = Math.max(calc_radius(container), 20);
        R = 20;
		r = Math.sqrt(R ** 2 - (R / 2) ** 2);
		sides = 6;

		text_size = R * 1.2;

		field = two.makeGroup();

		grid = [];
		initGrid(two, field, grid, R, R, sides, text_size, ticks);
		let id = 1;
		for (let round of rounds) {
			for (let move of round.attack.moves) {
				ticks[id] = {
					field: structuredClone(ticks[id - 1].field),
					statistics: structuredClone(ticks[id - 1].statistics)
				};
				ticks[id].field[move.from.row][move.from.col].value = move.from.value;
				ticks[id].field[move.to.row][move.to.col].value = move.to.value;
				if (move.is_winner) {
					ticks[id].field[move.to.row][move.to.col].color =
						ticks[id].field[move.from.row][move.from.col].color;
				}
				for (const name of names) {
					ticks[id].statistics.scores[name] = move.statistics.scores[name];
				}
				id++;
			}
			for (let cell of round.defence.cells) {
				ticks[id] = {
					field: structuredClone(ticks[id - 1].field),
					statistics: structuredClone(ticks[id - 1].statistics)
				};
				ticks[id].field[cell.row][cell.col].value = cell.value;
				for (const name of names) {
					ticks[id].statistics.scores[name] = round.defence.statistics.scores[name];
				}
				id++;
			}
		}
	});

	const wait = (ms: number) => new Promise((resolve) => setTimeout(resolve, ms));

	async function showTick(tick: Field, grid: GridCell[][], two: Two) {
		for (let i = 0; i < h_n; i++) {
			for (let j = 0; j < w_n; j++) {
				grid[i][j].inner.fill = get_cell_color(tick.field[i][j].color, 50, tick.field[i][j].value);
				grid[i][j].text.value = tick.field[i][j].value.toString();
				if (grid[i][j].text.value !== '0') {
					grid[i][j].text.fill = 'white';
					grid[i][j].text.stroke = 'black';
				} else {
					grid[i][j].text.fill = 'transparent';
					grid[i][j].text.stroke = 'transparent';
					grid[i][j].inner.fill = get_cell_color(0, 0, 0, 0);
				}
			}
		}
		two.update();
	}

	function initGrid(
		two: Two,
		field: Group,
		grid: GridCell[][],
		R: number,
		r: number,
		sides: number,
		text_size: number,
		fld: Field[]
	) {
		fld[0] = {
			field: [],
			statistics: {
				scores: {}
			}
		};

		for (let name of names) {
			fld[0].statistics.scores[name] = 0;
		}

		for (let i = 0; i < h_n; i++) {
			grid[i] = [];
			fld[0].field[i] = [];
			const y = padding + r + i * r + (i * padding) / 2;
			for (let j = 0; j < w_n; j++) {
				const x = padding + R + j * 3 * R + 2 * padding * j + (i % 2) * padding + R * 1.5 * (i % 2);
				if (!fieldMask[i * w_n + j]) {
					continue;
				}
				const group = two.makeGroup();
				const hex = two.makePolygon(0, 0, R, sides);
				hex.fill = 'hsla(0, 0%, 67%, 0.06)';
				hex.linewidth = 0;
				(hex as any).name = 'main-hex';

				const inner = two.makePolygon(0, 0, r * 0.9, sides);
				inner.rotation = Math.PI / 6;
				inner.fill = get_hsla_color(0, 6, 20, 1.0);
				inner.linewidth = 1;

				group.add(hex, inner);
				group.translation.set(x, y);

				const text = two.makeText('0', 0, 0);
				text.stroke = 'transparent';
				text.linewidth = 0.5;
				text.fill = 'transparent';
				text.size = text_size;
				text.weight = 800;
				text.family = '"Courier New", Courier, monospace';
				group.add(text);

				field.add(group);
				grid[i][j] = { group: group, hex: hex, inner: inner, text: text };
				fld[0].field[i][j] = {
					color: 0,
					position: {
						row: i,
						col: j
					},
					value: 0
				};
			}
		}

		for (let idx = 0; idx < start_positions.length; idx++) {
			const position = start_positions[idx];
			const ro = position.row;
			const co = position.col;
			grid[ro][co].inner.fill = get_cell_color(colors[idx], 50, 2);
			grid[ro][co].text.stroke = 'black';
			grid[ro][co].text.fill = 'white';
			grid[ro][co].text.value = '2';
			fld[0].field[ro][co].value = 2;
			fld[0].field[ro][co].color = colors[idx];
		}

		for (let cell of big_cells) {
			grid[cell.row][cell.col].inner.scale = (r + padding / 2) / r;
		}

		const bounds = field.getBoundingClientRect(true);
		field.translation.set(
			-bounds.left + (two.width - bounds.width) / 2,
			-bounds.top + (two.height - bounds.height) / 2
		);

		two.update();
		for (let i = 0; i < h_n; i++) {
			for (let j = 0; j < w_n; j++) {
				const element = grid[i][j].inner.renderer.elem as SVGElement;
				if (element) {
					element.classList.add('untouchable');
				}
			}
		}

		two.update();
	}

	function get_hsla_color(hue: number, percent: number, light: number, opacity = 1) {
		return `hsla(${hue}, ${percent}%, ${light}%, ${opacity})`;
	}

	function get_cell_color(hue: number, percent: number = 30, value: number, opacity: number = 1) {
		if (value !== 0) {
			const min_l = 20,
				max_l = 70;
			const max_value = 12;
			const ratio = Math.min(value / max_value, 1);
			const light = max_l - ratio * (max_l - min_l);
			return get_hsla_color(hue, 30, light);
		}
		return 'hsla(0, 6%, 20%, 1)';
	}

	function calc_radius(container: HTMLElement) {
		let le = 0,
			ri = Math.max(container.clientHeight, container.clientWidth);
		const safetyMargin = 20;
		while (le + 0.001 < ri) {
			let mid = (le + ri) / 2;
			const radius = Math.sqrt(mid ** 2 - (mid / 2) ** 2);
			const y =
				padding + radius + h_n * radius + (h_n * padding) / 2 + padding + mid * 2 + safetyMargin;
			const x =
				padding + mid + w_n * 3 * mid + 2 * padding * w_n + padding * 2 + mid * 2 + safetyMargin;
			if (x <= container.clientWidth && y <= container.clientHeight) {
				le = mid;
			} else {
				ri = mid;
			}
		}
		return Math.max(1, le - 2);
	}

	let inputValue = '0';

	function jumpToFrame() {
		const frameNum = parseInt(inputValue);
		if (!isNaN(frameNum) && frameNum >= 0 && frameNum < ticks.length) {
			currentTickIndex = frameNum;
		} else {
			inputValue = currentTickIndex.toString();
		}
	}

	function handleKeyPress(e: KeyboardEvent) {
		if (e.key === 'Enter') {
			jumpToFrame();
		}
	}

	$: inputValue = currentTickIndex.toString();

	let playbackSpeed = 1.0;
	let animationInterval: number | null = null;

	function togglePlayPause() {
		if (isPlaying) {
			pauseAnimation();
		} else {
			playAnimation();
		}
	}

	function playAnimation() {
		if (isPlaying) return;

		isPlaying = true;
		const interval = 1000 / playbackSpeed;

		animationInterval = window.setInterval(() => {
			if (currentTickIndex < ticks.length - 1) {
				currentTickIndex++;
			} else {
				pauseAnimation();
			}
		}, interval);
	}

	function pauseAnimation() {
		isPlaying = false;
		if (animationInterval !== null) {
			clearInterval(animationInterval);
			animationInterval = null;
		}
	}

	function increaseSpeed() {
		playbackSpeed = Math.min(playbackSpeed + 0.5, 5.0);
		if (isPlaying) {
			pauseAnimation();
			playAnimation();
		}
	}

	function decreaseSpeed() {
		playbackSpeed = Math.max(playbackSpeed - 0.5, 0.5);
		if (isPlaying) {
			pauseAnimation();
			playAnimation();
		}
	}

	function resetAnimation() {
		pauseAnimation();
		currentTickIndex = 0;
		playbackSpeed = 1.0;
	}

	import { onDestroy } from 'svelte';
	onDestroy(() => {
		if (animationInterval !== null) {
			clearInterval(animationInterval);
		}
	});
</script>

<div class="visualizer-outer">
	<div id="visualizer-container">
		<div class="canvas-wrapper">
			<div id="game-canvas" bind:this={container}></div>
		</div>

		<div class="controls-panel">
			<div class="slider-container">
				<input
					type="range"
					min="0"
					max={ticks.length - 1}
					bind:value={currentTickIndex}
					class="slider some-touchable"
					title="Перемотка по кадрам"
				/>
				<div class="frame-input-container">
					<input
						type="number"
						bind:value={inputValue}
						on:keypress={handleKeyPress}
						min="0"
						max={ticks.length - 1}
						class="frame-input some-touchable"
						title="Номер кадра"
					/>
					<span class="total-frames">/ {ticks.length - 1}</span>
				</div>
			</div>

			<div class="compact-controls">
				<div class="nav-buttons">
					<button
						on:click={() => (currentTickIndex = 0)}
						class="icon-btn some-button"
						title="В начало"
					>
						⏮️
					</button>
					<button
						on:click={() => (currentTickIndex > 0 ? currentTickIndex-- : null)}
						class="icon-btn some-button"
						title="Предыдущий кадр"
					>
						◀️
					</button>

					<button
						on:click={togglePlayPause}
						class="icon-btn play-btn some-button"
						title={isPlaying ? 'Пауза' : 'Воспроизвести'}
					>
						{#if isPlaying}
							⏸️
						{:else}
							▶️
						{/if}
					</button>

					<button
						on:click={() =>
							currentTickIndex < ticks.length - 1 ? currentTickIndex++ : currentTickIndex}
						class="icon-btn some-button"
						title="Следующий кадр"
					>
						▶️
					</button>
					<button
						on:click={() => (currentTickIndex = ticks.length - 1)}
						class="icon-btn some-button"
						title="В конец"
					>
						⏭️
					</button>
				</div>

				<div class="secondary-buttons">
					<button
						on:click={resetAnimation}
						class="icon-btn reset-btn some-button"
						title="Сбросить анимацию"
					>
						🔄
					</button>
					<button
						on:click={decreaseSpeed}
						class="icon-btn speed-btn some-button"
						title="Уменьшить скорость"
					>
						➖
					</button>
					<span class="speed-display">{playbackSpeed.toFixed(1)}x</span>
					<button
						on:click={increaseSpeed}
						class="icon-btn speed-btn some-button"
						title="Увеличить скорость"
					>
						➕
					</button>
				</div>
			</div>
		</div>
	</div>
</div>

<style>
	.visualizer-outer {
		display: flex;
		flex-direction: column;
		min-width: 1000px;
		min-height: 900px;
		width: 100%;
		height: calc(100% - 40px);
		position: relative;
		overflow: hidden;
	}
	#visualizer-container {
		display: flex;
		flex-direction: column;
		min-width: 1000px;
		min-height: 900px;
		width: 100%;
		height: 100%;
	}

	.canvas-wrapper {
		display: flex;
		align-items: center;
		justify-items: center;
		justify-content: center;
		flex: 1;
		position: absolute;
		overflow: hidden;
		min-height: 700px;
		height: 70%;
		width: 100%;
		background: rgba(50, 50, 50, 0.05);
		margin-top: 5%;
		margin-bottom: calc(max(10%, 200px));
	}

	#game-canvas {
		position: absolute;
		width: 100%;
		height: 100%;
		user-select: none;
		-webkit-user-select: none;
		cursor: default;
	}

	.controls-panel {
		position: fixed;
		bottom: 10px;
		left: 50%;
		transform: translateX(-50%);
		padding: 15px;
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 90%;
		max-width: 1200px;
		min-height: 120px;
		border-radius: 10px;
		box-sizing: border-box;
		margin: 0 auto;
		pointer-events: none;
		/* pointer-events: all; */
	}

	.some-button {
		pointer-events: all;
	}

	.some-touchable {
		pointer-events: all;
	}
	/* .controls-panel {
        position: fixed;
        bottom: 50px;
        margin: 0 auto;  
        padding: 5px;
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 90%;
        height: 120px;
        min-height: 120px;
    } */

	.slider-container {
		width: 100%;
		display: flex;
		align-items: center;
		gap: 5px;
	}

	.slider {
		flex: 1;
		height: 6px;
		-webkit-appearance: none;
		appearance: none;
		background: #667eea;
		cursor: pointer;
	}

	.slider::-webkit-slider-thumb {
		-webkit-appearance: none;
		appearance: none;
		width: 21px;
		height: 21px;
		background: white;
		border: 2px solid #667eea;
		cursor: pointer;
	}

	.slider::-moz-range-thumb {
		width: 21px;
		height: 21px;
		background: white;
		border: 2px solid #667eea;
		cursor: pointer;
	}

	.frame-input-container {
		display: flex;
		align-items: center;
		gap: 5px;
		min-width: 100px;
	}

	.frame-input {
		width: 60px;
		padding: 8px 12px;
		border: 1px solid rgba(255, 255, 255, 0.3);
		background: rgba(255, 255, 255, 0.1);
		border-radius: 6px;
		color: white;
		font-family: monospace;
		font-size: 15px;
		text-align: center;
		outline: none;
	}

	.frame-input:focus {
		border-color: #667eea;
		box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.3);
	}

	.total-frames {
		color: rgba(255, 255, 255, 0.7);
		font-family: monospace;
		font-size: 15px;
	}

	.compact-controls {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
		gap: 10px;
		width: 100%;
	}

	.nav-buttons,
	.secondary-buttons {
		display: flex;
		justify-content: center;
		align-items: center;
		gap: 5px;
		width: auto;
		min-height: 40px;
	}

	.icon-btn {
		width: 40px;
		height: 40px;
		padding: 0;
		border: none;
		color: white;
		cursor: pointer;
		font-size: 18px;
		display: flex;
		align-items: center;
		justify-content: center;
		transition:
			transform 0.2s ease,
			box-shadow 0.2s ease;
		backface-visibility: hidden;
		-webkit-backface-visibility: hidden;
		-webkit-font-smoothing: antialiased;
		text-rendering: geometricPrecision;
		-webkit-text-stroke: 0.01em transparent;
	}

	.icon-btn:hover {
		transform: translateY(-2px);
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
	}

	.icon-btn:active {
		transform: translateY(0);
	}

	.play-btn {
		width: 40px;
		height: 40px;
		font-size: 18px;
	}

	.speed-btn {
		width: 40px;
		height: 40px;
		font-size: 18px;
	}

	.speed-display {
		color: white;
		font-family: monospace;
		font-size: 15px;
		font-weight: bold;
		min-width: 50px;
		text-align: center;
		padding: 5px;
	}

	@media (max-width: 800px) {
		.controls-panel {
			margin: 10 auto;
			width: 95%;
			padding: 12px;
		}

		.slider-container {
			flex-direction: column;
			gap: 5px;
		}

		.frame-input-container {
			width: 100%;
			justify-content: center;
		}

		.compact-controls {
			display: flex;
			flex-direction: column;
			gap: 5px;
		}

		.nav-buttons,
		.secondary-buttons {
			gap: 5px;
			width: 100%;
			min-height: 70px;
		}

		.icon-btn {
			width: 45px;
			height: 45px;
			font-size: 20px;
		}

		.play-btn {
			width: 45px;
			height: 45px;
			font-size: 20px;
		}

		.speed-btn {
			width: 45px;
			height: 45px;
			font-size: 20px;
		}

		.speed-display {
			font-size: 13px;
			min-width: 40px;
		}
	}

	#visualizer-container {
		display: flex;
		flex-direction: column;
		min-width: 1000px;
		min-height: 800px;
		width: 100%;
		height: 100%;
	}

	#game-canvas {
		position: absolute;
		width: 100%;
		height: 100%;
		user-select: none;
		overflow: hidden;
		-webkit-user-select: none;
		cursor: default;
	}

	.frame-input {
		-moz-appearance: textfield;
		appearance: textfield;
	}

	.frame-input::-webkit-outer-spin-button,
	.frame-input::-webkit-inner-spin-button {
		-webkit-appearance: none;
		margin: 0;
	}

	.frame-input[type='number'] {
		-moz-appearance: textfield;
	}

	.slider-container {
		width: 100%;
		display: flex;
		align-items: center;
		gap: 15px;
	}

	.slider {
		flex: 1;
		height: 6px;
		-webkit-appearance: none;
		appearance: none;
		background: #667eea;
		border-radius: 3px;
		cursor: pointer;
	}

	.slider::-webkit-slider-thumb {
		-webkit-appearance: none;
		appearance: none;
		width: 20px;
		height: 20px;
		border-radius: 50%;
		background: white;
		border: 2px solid #667eea;
		cursor: pointer;
	}

	.slider::-moz-range-thumb {
		width: 20px;
		height: 20px;
		border-radius: 50%;
		background: white;
		border: 2px solid #667eea;
		cursor: pointer;
	}

	.frame-input-container {
		display: flex;
		align-items: center;
		gap: 5px;
		min-width: 100px;
	}

	.frame-input {
		width: 60px;
		padding: 8px 12px;
		border: 1px solid rgba(255, 255, 255, 0.3);
		background: rgba(255, 255, 255, 0.1);
		border-radius: 6px;
		color: white;
		font-family: monospace;
		font-size: 15px;
		text-align: center;
		outline: none;
	}

	.frame-input:focus {
		border-color: #667eea;
		box-shadow: 0 0 0 2px rgba(102, 126, 234, 0.3);
	}

	.total-frames {
		color: rgba(255, 255, 255, 0.7);
		font-family: monospace;
		font-size: 15px;
	}

	#game-canvas {
		position: absolute;
		width: 100%;
		height: 100%;
		min-height: 0;
		overflow: hidden;
		margin-bottom: 1%;

		user-select: none;
		-webkit-user-select: none;
		cursor: default;
	}
</style>
