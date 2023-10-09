<script lang="ts">
    import Header from "../../components/Header.svelte";
    import Login from "./Login.svelte";
</script>

<div class="app">
    <Header />

    <div class="account-wrapper">
        <Login />   
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

    .account-wrapper {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100%;
        width: 100%;
    }

</style>