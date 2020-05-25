<style>
  .userbutton {
    margin-right: 1em;
  }

  .btn-small {
    color: #eee;
  }

  .img {
    width: 48px;
  }
</style>

<script>
  import { getAuth } from "./auth.js";

  const {
    isLoading,
    isAuthenticated,
    login,
    logout,
    authToken,
    authError,
    userInfo,
  } = getAuth();

  async function loadUserData() {
    axios(`http://localhost:8010/api/user`, {
      method: "GET",
      mode: "no-cors",
      credentials: "same-origin",
      headers: {
        Authorization: `Bearer ${$authToken}`,
      },
    })
      .then((result) => console.log(result))
      .catch((err) => console.log(err));
  }

  async function signup() {
    await login();

    if ($isAuthenticated) {
      await loadUserData();
    }
  }
</script>

<!-- Dropdown Structure -->
<ul id="dropdown1" class="dropdown-content">
  <li>
    <a href="#!">one</a>
  </li>
  <li>
    <a href="#!">two</a>
  </li>
  <li class="divider"></li>
  <li>
    <a href="#!">three</a>
  </li>
</ul>

{#if $isLoading}
  <li class="right-align">
    <p>Loading ...</p>
  </li>
{:else if $authError}
  <li class="right-align">
    <p>Got error: {$authError.message}</p>
  </li>
{:else if !$isAuthenticated}
  <li class="right-align">
    <p on:click="{() => signup()}" class="btn-small userbutton green">Signup</p>
  </li>
  <li class="right-align">
    <button on:click="{() => login()}" class="btn-small userbutton green">
      Log in
    </button>
  </li>
{:else}
  <li class="right-align">
    <button on:click="{() => logout()}" class="btn-flat">
      Log out
    </button>
  </li>
  <li>
    <a class="dropdown-trigger" href="#!" data-target="dropdown1">
      {$userInfo.name}
      <i class="material-icons right">arrow_drop_down</i>
    </a>
  </li>
  <li>
    <img src="{$userInfo.picture}" alt="" class="circle img " />

  </li>
{/if}
