<script lang="ts">
    import { page } from '$app/state';
    import Visualizer from '$lib/components/Visualizer.svelte';

    const number = page.params.number;


    export let data;
    interface Position {
        row: number,
        col: number,
    }
    interface Cell {
        row: number,
        col: number,
        value: number,
    };

    interface AttackMove {
        from: Cell,
        to: Cell,
        is_winner: boolean,
        statistics: Statistics,
    }

    interface Attack {
        moves: AttackMove[],
    }

    interface Defence {
        cells: Cell[],
        statistics: Statistics,
    }

    interface Round {
        attack: Attack;
        defence: Defence;
    };

    interface Statistics {
        scores: Record<string, number>,
    }

    const nrows: number = data.data.map.nrows;
    const ncols: number = data.data.map.ncols;
    const field: boolean[] = data.data.map.fieldMask;
    const start_positions: Position[] = data.data.start_positions;
    const big_cells: Position[] = data.data.map.bigCells;
    const rounds: Round[] = data.data.rounds;
    const names: string[] = data.data.names;
    const colors: number[] = [];
    const colorMap: Map<string, string> = new Map();
    
    import { onMount } from 'svelte';

    onMount(() => {
        let resizeTimer: number;

        window.addEventListener('resize', () => {
            clearTimeout(resizeTimer);
            
            resizeTimer = setTimeout(() => {
                location.reload();
            }, 250);
        });
        for (let i = 0; i < names.length; i++) {
            colors.push(360 / (names.length + 2) * (i + 1));
        }
        for (let i = 0; i < names.length; i++) {
            colorMap.set(names[i], `hsla(${colors[i]}, 50%, 60%, 1)`)
        }

        
    });
    let currentScores: Record<string, number> = {};
</script>

<svelte:head>
    <title>Симуляция {number}</title>
    
</svelte:head>

<div id="my-main-page-container" class="page-container">
    <ul class="overlay-list">
        {#each names as name, i}
            <li class="overlay-list-item" style="font-weight: 900; -webkit-text-stroke: 0.7px black; color: hsla({360 / (names.length + 2) * (i + 1)}, 50%, 60%, 1)">
                {name}: {currentScores[name] ?? 0}
            </li>
        {/each}
    </ul>
    <div class="visualizer-wrapper">
        <Visualizer 
            bind:currentScores={currentScores}
            n={nrows} 
            m={ncols} 
            fieldMask={field} 
            start_positions={start_positions} 
            big_cells={big_cells} 
            rounds={rounds}
            names={names}
        />
    </div>
</div>

<style>
    .overlay-list {
        position: absolute;
        pointer-events: none;
        z-index: 1000;

        user-select: none;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
    }

    .overlay-list-item {
        margin-left: 10px;
        margin-top: 10px;
        font-size: 4vh;
        font-family: "Courier New", Courier, monospace;
    }
    .page-container {
        position: absolute;
        width: 100%;
        min-height: 100vh;
        padding: 20px;
        overflow: hidden;
        box-sizing: border-box;
    }

    .visualizer-wrapper {
        scrollbar-width: none;
        
        -ms-overflow-style: none;
    }

    .visualizer-wrapper::-webkit-scrollbar {
        display: none;
    }

    .visualizer-wrapper {
        position: absolute;
        z-index: 1;
        width: calc(100% - 40px);
        overflow: auto;
        height: calc(100% - 40px);
        background-color: var(--primary-bg-color);
        border-radius: 12px;
        box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
        display: flex;
        flex-direction: column;
    }
</style>