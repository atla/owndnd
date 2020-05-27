<style>
  #terminalWindow {
    padding: 1em;
    background: #000;
    border-width: 2px;
    border-style: solid;
    border-color: #33ff2266;
    border-radius: 0.5em;
  }
</style>

<script>
  import "../../node_modules/xterm/css/xterm.css";
  import { onMount, onDestroy } from "svelte";
  import { createAuth, getAuth } from "../auth.js";
  import axios from "axios";
  import xterm from "xterm";
  import LocalEchoController from "./echo/LocalEchoController";
  import fit from "xterm-addon-fit";
  import { createClient, getClient } from "./Client";

  let client;

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
  };

  $: {
  }
  function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  function readLine(localEcho, term) {
    localEcho
      .read("~$ ")
      .then((input) => {
        client.onInput(input);
        readLine(localEcho, term);
      })
      .catch((error) => console.log(`Error reading: ${error}`));
  }

  const createRenderer = (term) => {
    return (data) => {
      term.writeln(data);
    };
  };

  async function setupTerminal() {
    var term = new xterm.Terminal();
    var fitAddon = new fit.FitAddon();

    term.loadAddon(fitAddon);

    // setup terminal
    term.setOption("cursorBlink", true);
    term.setOption("convertEol", true);

    term.open(document.getElementById("terminal"));
    fitAddon.fit();

    const localEcho = new LocalEchoController(term);
    localEcho.addAutocompleteHandler(autocompleteCommonCommands);
    term.writeln("Connected to [Tales of the Red Dragon's Lair] ...");

    client = createClient(createRenderer(term));

    readLine(localEcho, term);
  }

  onMount(async () => {
    // change global background
    document.body.style.backgroundImage = "url('bg.jpg')";
    document.body.style.backdropFilter =
      "blur(10px) saturate(30%) brightness(50%)";
    setupTerminal();
  });

  onDestroy(async () => {
    // change global background
    document.body.style.backgroundImage = "";
    document.body.style.backdropFilter = "";
    setupTerminal();
  });

  function autocompleteCommonCommands(index, tokens) {
    if (index == 0) return ["north", "east", "south", "west", "say"];
    return [];
  }
</script>

<div id="terminalWindow" class="z-depth-3">
  <div id="terminal"></div>
</div>
