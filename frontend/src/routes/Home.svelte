<script lang="ts">
  import { createGame, joinGame } from '../store/game';
  import { toggleModal } from '../store/modal';

  let input = "";

  function submitJoinRequest(id: string) {
    if (!id.length) {
      toggleModal({
        title: 'Error!',
        description: 'Please enter a valid game ID!',
        status: 'error'
      });

      return;
    }

    return joinGame(id);
  }
</script>

<main>
  <h1>Welcome to tic-tac-toe!</h1>

  <button on:click={createGame}>Create game</button>
  <button on:click={() => submitJoinRequest(input)}>Join game</button>

  <input id="input" type="text" bind:value={input}>
</main>