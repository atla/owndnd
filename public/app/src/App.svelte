<style>
  :global(body) {
    background-color: #232333;
    color: #d8dee9;
    transition: background-color 0.3s;
    margin: 0 auto;
    padding: 0px;

    /* The image used */
    /*background-image: url("bg.jpg");*/
    /* Full height */
    height: 100%;
    /* Center and scale the image nicely */
    background-position: center;
    background-repeat: no-repeat;
    background-size: cover;

  }
  .userbutton {
    margin-left: 1em;
    margin-right: 1em;
  }
  :global(a:visited) {
    text-decoration: none;
    text-decoration-line: none;
  }
  nav {
    background: #00000033;
  }
  nav {
    margin-bottom: 2em;
  }

  img {
    height: 40px;
  }
</style>

<script>
  import { Router, Link, Route } from "svelte-routing";
  import { afterUpdate, onMount } from "svelte";

  import { user } from "./stores.js";
  import { fade } from "svelte/transition";

  import Game from "./game/Game.svelte";

  import UserMenu from "./UserMenu.svelte";
  import Input from "./Input.svelte";
  import Home from "./Home.svelte";
  import Welcome from "./Welcome.svelte";
  import Creator from "./creator/Creator.svelte";
  import UserForm from "./UserForm.svelte";
  import { createAuth } from "./auth.js";

  // Auth0 config
  const config = {
    domain: "owndnd.eu.auth0.com",
    client_id: "mxcEqTuAUOzrL798mbVTpqFxpGGVp3gI",
    audience: "http://talesofapirate.com/dnd/api",
  };

  const {
    isLoading,
    isAuthenticated,
    login,
    logout,
    authToken,
    authError,
    userInfo,
  } = createAuth(config);

  $: state = {
    isLoading: $isLoading,
    isAuthenticated: $isAuthenticated,
    authError: $authError,
    userInfo: $userInfo ? $userInfo.name : null,
    authToken: $authToken.slice(0, 20),
  };

  onMount(async () => {
    //var elems = document.querySelectorAll(".tabs");
    //let instance = M.Tabs.init(elems);
  });
</script>

<svelte:head>
  <script src="https://cdn.auth0.com/js/auth0/9.0/auth0.min.js">

  </script>
  <link
    rel="stylesheet"
    href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css"
  />
  <script
    src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0-beta/js/materialize.min.js">

  </script>
  <link
    rel="stylesheet"
    href="https://fonts.googleapis.com/icon?family=Material+Icons"
  />

</svelte:head>

<Router>
  <nav class="nav-extended">
    <div class="nav-wrapper container">
      <a href="#" class="brand-logo">
        <span class="valign-wrapper italic">
          <i class="small material-icons prefix">menu_book</i>
          <Link to="/">Tales</Link>
        </span>
      </a>

      <ul id="nav-mobile" class="right hide-on-med-and-down">
        <li class="align-left">
          <Link to="play">
            <button on:click class="btn-small userbutton blue">Play</button>
          </Link>

        </li>
        {#if $isAuthenticated}
          <li>
            <Link to="creator">Creator</Link>
          </li>
        {/if}
        <li>
          <Link to="signup">News</Link>
        </li>
        <UserMenu />
      </ul>

    </div>
    <div class="nav-content container">
      <ul class="tabs tabs-transparent">
        <li class="tab">
          <a href="#test1">Rooms</a>
        </li>
        <li class="tab">
          <a href="#test2">NPCs</a>
        </li>
        <li class="tab">
          <a href="#test3">Items</a>
        </li>
        <li class="tab">
          <a href="#test4">Quests</a>
        </li>
        <li class="tab">
          <a href="#world">World</a>
        </li>
      </ul>
    </div>
  </nav>

  <main class="container">

    <Route path="account" component="{UserForm}" />
    <Route path="creator" component="{Creator}" />
    <Route path="play" component="{Game}" />

    <Route path="/">
      {#if $isAuthenticated}
        <UserForm />
      {:else}
        <Welcome />
      {/if}
    </Route>
  </main>
</Router>
