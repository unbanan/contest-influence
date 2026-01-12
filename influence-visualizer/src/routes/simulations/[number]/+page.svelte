<script>
    import { page } from '$app/state';
    const number = page.params.number;
    
    export let data;
    console.log(data)

</script>

<svelte:head>
    <title>Симуляция {number}</title>
    
</svelte:head>
