<script lang="ts">
  import { Router, Route, navigate } from "svelte-routing";
  import { gameStore } from "./store/game";
  import Game from "./routes/Game.svelte";
  import Home from "./routes/Home.svelte";
  import Modal from "./lib/Modal.svelte";

  export let url = "";

  gameStore.subscribe((store) => {
    if (store.id) {
      console.log(store);
      navigate('/game');
    } else {
      navigate('/');
    }
  });
</script>

<Router {url}>
  <Route path="/game">
    <Game />
  </Route>
  <Route path="/">
    <Home />
  </Route>

  <Modal />
</Router>
