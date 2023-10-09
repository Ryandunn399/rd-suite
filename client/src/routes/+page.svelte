<script lang="ts">
    import Homepage from "./Homepage.svelte";
    import Header from "../components/Header.svelte";
</script>


<div class="app">
    <Header />

    <div class="home">
        <Homepage />
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