<script lang="ts">
    import { redirectLocal } from "../utilities/utilites";
</script>

<div class="navbar">
    <div class="nav-container empty-container"></div>
    <div class="nav-container title">TOOL SUITE</div>
    <div class="nav-container">
        <div class="login-icon">
            <button class="btn-login material-symbols-outlined md-2p5">account_circle</button>
        </div>
    </div>
</div>

<style>
    @import 'material-symbols';

    .navbar {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        width:100%;

        -moz-box-sizing: border-box; 
        -webkit-box-sizing: border-box; 
        box-sizing: border-box; 
    }

    .title {
        font-size: 3rem;
        font-family: var(--font-cubano);
        color: var(--active-color);
    }

    .nav-container {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 100%;
        height: 100%;

        -moz-box-sizing: border-box; 
        -webkit-box-sizing: border-box; 
        box-sizing: border-box; 
    }

    .login-icon {
        display: flex;
        justify-content: end;
        align-items: center;
        padding-right: 2rem;
        width: 100%;
    }

    .material-symbols-outlined.md-2p5 {
        font-size: 2.5rem;
    }

    .btn-login {
        padding: 0;
        border: none;
        background: none;
    }

    .btn-login:hover {
        color: var(--active-color);
        cursor: pointer;
    }

</style>