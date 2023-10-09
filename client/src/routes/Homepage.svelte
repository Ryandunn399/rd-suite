<script lang="ts">
    import Github from '$lib/images/github.png'
    import DND from '$lib/images/dnd.png'
    import FVTT from '$lib/images/fvtt.png'
    import Settings from '$lib/images/settings.png'

    import { redirectExternal, redirectLocal } from '../utilities/utilites'

</script>

<div class="row">
    <button class="btn-git" on:click={redirectExternal("github.com/ryandunn399")}>
        <img src={Github} alt="Github" class="logo git">
    </button>

    <button class="btn-fvtt">
        <img src={FVTT} alt="Foundry" class="logo fvtt">
    </button>
</div>

<div class="row">
    <button class="btn-dnd">
        <img src={DND} alt="Dungeons and Dragons" class="logo dnd">
    </button>

    <button class="btn-settings" on:click={redirectLocal("configs")}>
        <img src={Settings} alt="Settings" class="logo settings">
    </button>
</div>

<style lang="scss">
    @import '$lib/animations.scss';

    .row {
        display: flex;
        align-items: center;
        justify-content: center;
    }

    button {
        margin: 1rem;
        height: 12rem;
        width: 12rem;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: var(--background-secondary);
        box-shadow: rgba(0, 0, 0, 0.25) 0px 14px 28px, rgba(0, 0, 0, 0.22) 0px 10px 10px;
        transition: transform 330ms ease-in-out;
    }

    img {
        height: auto;
        width: 75%;
    }

    button:hover {
        cursor: pointer;
        transform: scale(1.1);
    }

    button.btn-git:hover {
        
        img {
            filter: drop-shadow(0 0 2em black);
        }

    }

    button.btn-dnd:hover {
        
        img {
            filter: drop-shadow(0 0 2em #ff3e00aa);
        }

    }

    button.btn-fvtt:hover {
        
        img {
            filter: drop-shadow(0 0 2em var(--active-color));
        }
    }

    button.btn-settings:hover {
        
        img {
            filter: drop-shadow(0 0 2em var(--active-color));
        }
    }

</style>