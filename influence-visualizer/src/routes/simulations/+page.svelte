<script>
    const sim_number = 100;
    const sims = Array.from({ length: sim_number }, (_, i) => i);
</script>

<svelte:head>
    <title>
        Список симуляций
    </title>   
</svelte:head>

<div class="ul-container">
    <ul class="ul-simulations">
        {#each sims as sim}
            <li class="li-simulations">
                <a class="a-simulations" href="/simulations/{sim}">
                    Симуляция №{sim}
                </a>
            </li>
        {/each}
    </ul>   
</div> 

<style>
    :root {
        --bg-color-1: #ababab0f;
        --bg-color-2: #353b48;
        --debug-color: pink;
        --link-color-1: #fbc531;
        --scrollbar-color: #fbc531;
        --link-color-2: #e9e8e5;
        --scrollbar-width: 10px;
        --scrollbar-height: 10px;
    }
    .ul-container {
        margin-top: 1vh;
        margin-bottom: 1vh;
        padding-top: 1vh;
        padding-bottom: 1vh;
        padding-right: 1vw;;
        padding-left: 1vw;
        position: relative;
        left: 15vw;
        width: 70vw;
        height: 98vh;
        background-color: var(--bg-color-1);
    }
    .ul-simulations {
        max-height: 96vh;
        width: calc(69vw - var(--scrollbar-width));
        overflow-x: hidden;
        overflow-y: scroll;
    }
    .a-simulations {
        font-size: 1.3rem;
        font-weight: 600;
        font-family: 'Courier New', Courier, monospace;
        color: var(--link-color-1);
        transition: 0.5s ease-out;
    }

    .a-simulations:hover {
        color: var(--link-color-2);
    }


    .ul-simulations::-webkit-scrollbar {
        width: var(--scrollbar-width);
        height: var(--scrollbar-height);
    }

    .ul-simulations::-webkit-scrollbar-track {
        background: var(--bg-color-2);
        border-radius: 0;
    }

    .ul-simulations::-webkit-scrollbar-thumb {
        background: var(--scrollbar-color);
        border-radius: 0;
        border: 3px solid --bg-color-2;
    }

    .ul-simulations::-webkit-scrollbar-thumb:hover {
        background: var(--scrollbar-color);
    }

    .ul-simulations::-webkit-scrollbar-corner {
        background: var(--bg-color-2);
    }
</style>