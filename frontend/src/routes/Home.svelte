<script lang="ts">
  import Modal from '../lib/Modal.svelte';
  import { toggleModal } from '../store';

  async function createGame() {
    try {
      const res = await fetch('http://localhost:8000/v1/game/create', {
        method: 'POST',
      });
  
      const data = await res.json();

    } catch (e) {
      toggleModal({
        title: 'Error!',
        description: 'Something went wrong!',
        status: 'error'
      });
    }
  }

  async function joinGame() {

  }
</script>

<main>
  <h1>Welcome to tic-tac-toe!</h1>

  <button on:click={createGame}>Create game</button>
  <button on:click={joinGame}>Join game</button>
</main>