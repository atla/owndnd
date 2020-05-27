// src/auth.js

import { onMount, setContext, getContext } from "svelte";
import { writable } from "svelte/store";

const room = {
  id: "Dungeon001_Room1",
  name: "Main Chamber",
  description:
    "You reach the Main Chamber of the Catacomb. The noise increases but you can't make out the origin of it.",
  detail:
    "You look closer to all sides of the room. After a thorough investigation you can see that parts of a wall are made up of loose rocks. You might be able to [move] these rocks.",
  exits: [
    {
      exit: "north",
      description: "Follow the door to the left",
      target: "Dungeon001_Entrance",
    },
    {
      exit: "hidden path",
      hidden: true,
      description: "You follow the hidden path on the east wall",
      target: "Dungeon001_End",
    },
  ],
  actions: [
    {
      action: "move rocks",
      description:
        "You try to move one of the medium sized rocks. Parts of the wall start to crumble and a hidden path opens up.",
    },
  ],
};

//const isLoading = writable(true);
const GAME_CLIENT = {};

function createClient(renderer) {
  const onInput = async (data) => {
    console.log(`client received data: ${data}`);

    msg = handleInput(data);

    renderer(msg);
  };

  const renderRoom = async () => {
    renderer(room.description);
  };

  const handleInput = async (data) => {
    return `### ${data}`;
  };

  const client = {
    onInput,
  };

  renderRoom();

  // setInterval(function () {
  //   renderer("\n<The lights in front of you are flickering>")
  // }, 5000);

  setContext(GAME_CLIENT, client);
  return client;
}

function getClient() {
  return getContext(AUTH_KEY);
}

export { createClient, getClient };
