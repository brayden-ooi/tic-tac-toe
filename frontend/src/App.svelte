<script lang="ts">
  import { Router, Route, navigate } from "svelte-routing";
  import { INITIAL_STATE, store } from './store';
  import Game from "./routes/Game.svelte";
  import Home from "./routes/Home.svelte";
  import Modal from "./lib/Modal.svelte";

  export let url = "";

  store.subscribe((store) => {
    if (store.game.id) {
      console.log(store.game);
      navigate('/game');
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
