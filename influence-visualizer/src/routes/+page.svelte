<svelte:head>
    <title>Визуализатор</title>
</svelte:head>

<div id="my-page-container">
    <div id="venom-container">
        <div id="venom-text">
            Venom
        </div>
        <a id="simulation-link" href="/simulations">
            Симуляции
        </a>
    </div>
</div>

<style>
    :root {
        --bg-color-1: #ababab0f;
        --bg-color-2: #353b48;
        --debug-color: pink;
        --link-color-1: #fbc531;
        --scrollbar-color: #fbc531;
        --font-color: rgb(226, 222, 222);
        --link-color-2: #e9e8e5;
        --scrollbar-width: 10px;
        --scrollbar-height: 10px;
    }
    #my-page-container {
        position: relative;
        width: 100vw;
        height: 100vh;
    }
    #venom-container {
        position: relative;
        top: 20vh;
        width: 100vw; 
        height: 18vh;
    }
    #venom-text {
        height: 10vh; 
        justify-content: center; 
        align-items: center;
        position: relative; 
        display: flex; 
        font-size: 10vh; 
        font-family: 'Courier New', Courier, monospace; 
        color: var(--font-color);
    }
    #simulation-link {
        /* background-color: red; */
        cursor: pointer; 
        height: 8vh; 
        justify-content: center; 
        align-items: center;
        position: relative; 
        display: flex; 
        font-size: 6vh; 
        font-family: 'Courier New', Courier, monospace;
        width: fit-content;
        margin: 0 auto;
        font-weight: 600;
        color: var(--link-color-1);
        transition: 0.5s ease-out;
    }
    #simulation-link:hover {
        color: var(--link-color-2);
    }

</style>