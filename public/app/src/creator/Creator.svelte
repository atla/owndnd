<style>

</style>

<script>
  import { writable } from "svelte/store";
  import { onMount } from "svelte";
  import { createAuth, getAuth } from "../auth.js";

  import axios from "axios";
  import {
    getRoom,
    getRooms,
    updateRoom,
  } from "../api/rooms.js";

  let rooms = [];

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
    rooms,
  };

  $: {
    getRooms($authToken, (data) => {
      rooms = data;
    });
  }

  onMount(async () => {
    var elems = document.querySelectorAll(".tabs");
    let instance = M.Tabs.init(elems);

    document.body.style.backgroundImage = "";
  });
</script>

<div class="row">
  <h4>Manage Rooms</h4>

  <table>
    <thead>
      <tr>
        <th>Id</th>
        <th>Name</th>
        <th>Description</th>
        <th>Detail</th>
      </tr>
    </thead>
    <tbody>
      {#each rooms as room}
        <tr>
          <td>{room.id}</td>
          <td>{room.name}</td>
          <td>{room.description}</td>
          <td>{room.detail}</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
