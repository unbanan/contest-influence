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
    }

    interface Attack {
        moves: AttackMove[],
    }

    interface Defence {
        cells: Cell[],
    }

    interface Round {
        attack: Attack;
        defence: Defence;
    };
    const nrows: number = data.data.map.nrows;
    const ncols: number = data.data.map.ncols;
    const field: boolean[] = data.data.map.fieldMask;
    const start_positions: Position[] = data.data.start_positions;
    const big_cells: Position[] = data.data.map.bigCells;
    const rounds: Round[] = data.data.rounds;
</script>

<svelte:head>
    <title>Симуляция {number}</title>
    
</svelte:head>

<div class="page-container">
    <div class="visualizer-wrapper">
        <Visualizer 
            n={nrows} 
            m={ncols} 
            fieldMask={field} 
            start_positions={start_positions} 
            big_cells={big_cells} 
            rounds={rounds}
        />
    </div>
</div>

<style>
    .page-container {
        width: 100%;
        height: 100vh;
        padding: 20px;
        box-sizing: border-box;
    }

    .visualizer-wrapper {
        width: 100%;
        height: 100%;
        background-color: var(--primary-bg-color);
        border-radius: 12px;
        box-shadow: 0 4px 20px rgba(0, 0, 0, 0.15);
        overflow: auto;
        display: flex;
        flex-direction: column;
    }
</style>