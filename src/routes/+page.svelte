<script lang="ts">
    import App from "./App.svelte";
    import Header from "./Header.svelte";
</script>


<div class="app">
    <Header />

    <div class="home">
        <App />
    </div>
</div>

<style>

    .app {
        display: flex;
        flex-direction: column;
        height: 99vh;
        width: 100%;
        align-items: center;
    }

    .home {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100%;
        width: 100%;
    }
</style>