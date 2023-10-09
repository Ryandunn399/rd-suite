<script lang="ts">
    import { useForm, minLength, validators, required } from 'svelte-use-form';

    let username: string;
    let password: string;

    const form = useForm();

    function handleSubmit(e: Event) {
        e.preventDefault();
        console.log(username + " " + password);
        username = "";
        password = "";
    }
</script>

<form class="login-form" use:form on:submit={handleSubmit} method="post" action="?/login">
    
    <div class="login-wrapper">
        
        <div class="input-wrapper">
            <label class="input-label" for="user">Username</label>
            <input 
                type="text" 
                name="user" 
                use:validators={[required]}
                bind:value={username}    
            />
        </div>
    
        <div class="input-wrapper">
            <label class="input-label" for="pass">Password</label>
            <input
                type="password" 
                name="pass"
                use:validators={[required]}
                bind:value={password}
            />
        </div>

        <div class="input-wrapper">
            <button class="submit-btn">Submit</button>
        </div>
    </div>
</form>

<style>

    .login-form {
        display: flex;
        flex-direction: column;
        background-color: var(--background-secondary);
        border-radius: 0.5rem;
        width: 24%;
    }

    .input-wrapper {
        display:flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        padding-top: 5%;
    }

    .input-label {
        width: 61%;
        padding-bottom: 0.5%;
        font-family: var(--font-mono);
    }

    input {
        height: 2rem;
        width: 60%;
    }

    .submit-btn {
        background-color: var(--active-color);
        font-family: var(--font-mono);
        font-weight: 800;
        color: white;
        font-size: 1.5rem;
        margin-top: 5%;
        margin-bottom: 5%;
    }

</style>