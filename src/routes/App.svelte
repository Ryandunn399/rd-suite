<script lang="ts">
    import Github from '$lib/images/github.png'
    import DND from '$lib/images/dnd.png'
    import FVTT from '$lib/images/fvtt.png'
    import Settings from '$lib/images/settings.png'


    const redirectLocal = (location: string) => () => {
        window.location.href = "/" + location;
    }

    const redirectExternal = (location: string) => () => {
        window.location.href = "https://" + location;
    }

</script>

<div class="row">
    <button on:click={redirectExternal("github.com/ryandunn399")}>
        <img src={Github} alt="Github" class="logo git">
    </button>

    <button>
        <img src={FVTT} alt="Foundry" class="logo fvtt">
    </button>
</div>

<div class="row">
    <button>
        <img src={DND} alt="Dungeons and Dragons" class="logo dnd">
    </button>

    <button on:click={redirectLocal("configs")}>
        <img src={Settings} alt="Settings" class="logo settings">
    </button>
</div>

<style lang="scss">
    @import '$lib/animations.scss';

    .row {
        display: flex;
        margin: 1rem;
        align-items: center;
        justify-content: center;
    }

    button {
        margin: 0rem 1rem;
        height: 12rem;
        width: 12rem;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: var(--background-secondary);
        box-shadow: rgba(0, 0, 0, 0.25) 0px 14px 28px, rgba(0, 0, 0, 0.22) 0px 10px 10px;
        @include magic-border(4px, var(--active-color), 0.2s, 1);
    }

    img {
        height: auto;
        width: 75%;
    }

    .button:hover {
        cursor: pointer;
    }

    .logo.git:hover {
        filter: drop-shadow(0 0 2em black);
    }

    .logo.dnd:hover {
        filter: drop-shadow(0 0 2em #ff3e00aa);
    }

    .logo.fvtt:hover {
        filter: drop-shadow(0 0 2em var(--active-color));
    }

    .logo.settings:hover {
        filter: drop-shadow(0 0 2em var(--active-color));
    }

</style>