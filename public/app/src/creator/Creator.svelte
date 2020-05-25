<style>

</style>

<script>
  import { writable } from "svelte/store";
  import { onMount } from "svelte";
  import { createAuth, getAuth } from "../auth.js";

  import axios from "axios";

  let characters = [];

  const {
    isLoading,
    isAuthenticated,
    login,
    logout,
    authToken,
    authError,
    userInfo,
  } = getAuth();

  $: state = {
    isLoading: $isLoading,
    isAuthenticated: $isAuthenticated,
    authError: $authError,
    userInfo: $userInfo ? $userInfo.name : null,
    authToken: $authToken.slice(0, 20),
    characters,
  };

  $: {
    axios(`http://localhost:8010/api/characters`, {
      method: "GET",
      mode: "no-cors",
      credentials: "same-origin",
      headers: {
        Authorization: `Bearer ${$authToken}`,
      },
    }).then((r) => (characters = r.data));
  }

  onMount(async () => {
    var elems = document.querySelectorAll(".tabs");
    let instance = M.Tabs.init(elems);
  });
</script>

<div class="row">
  <h4>Manage Rooms</h4>

  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Description</th>
        <th>Created</th>
      </tr>
    </thead>
    <tbody>
      {#each characters as character}
        <tr>
          <td>{character.name}</td>
          <td>{character.description}</td>
          <td>{character.created}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
