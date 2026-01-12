<script lang="ts">
    import { page } from '$app/state';
    const number = page.params.number;

    import { onMount, onDestroy } from 'svelte';
    import Two from 'two.js';

    class Cell {
		x: number;
		y: number;
		value: number;
		color: string;
		isBig: boolean;
		label: string;
        constructor(x: number, y: number, value: number, color: string, isBig: boolean, label: string) {
            this.x = x;
            this.y = y;
            this.value = value;
            this.color = color;
            this.isBig = isBig;
            this.label = label;
        }
    }

    class Player {
        name: string;
        color: string;
        constructor(name: string, color: string) {
            this.name = name;
            this.color = color;
        }
    }

    class Field {
        cells: Cell[][];
        constructor(cells: Cell[][]) {
            this.cells = cells.map(row =>
                row.map(cell => ({ ...cell }))
            )
        }
    }

    class Tick {
        field: Field;
        turn: number;
        constructor(field: Field, turn: number) {
            this.field = new Field(field.cells);
            this.turn = turn;
        }
    }

</script>

<svelte:head>
    <title>Симуляция {number}</title>
    
</svelte:head>


<div id="game-canvas">
</div>