<svelte:head>
    <title>
        Страница не найдена
    </title>
</svelte:head>

<div id="my-page-container">
    <div id="my-error-message">
        <div id="title-message">
            Страница не найдена
        </div>
        <a id="main-page-link" href="/">
            Главная страница
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
    #my-error-message {
        /* background-color: var(--debug-color); */
        position: relative;
        height: 18vh;
        top: 20vh;
        width: 100vw;
    }
    #title-message {
        /* background-color: red; */
        position: relative;
        width: 100vw;
        height: 10vh;
        font-size: 10vh;
        font-family: 'Courier New', Courier, monospace;
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--font-color);
        font-weight: 300;
    }

    #main-page-link {
        position: relative; 
        font-size: 8vh; 
        font-family: 'Courier New', Courier, monospace;
        font-weight: 600;
        cursor: pointer; 
        height: 8vh; 
        display: flex; 
        justify-content: center; 
        width: fit-content;
        margin: 0 auto;
        align-items: center;
        color: var(--link-color-1);
        transition: 0.5s ease-out;
    }

    #main-page-link:hover {
        color: var(--link-color-2);
    }
</style>