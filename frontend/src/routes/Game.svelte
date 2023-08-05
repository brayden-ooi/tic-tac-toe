<script lang="ts">
  import Board from "../lib/Board.svelte";
  import { INITIAL_STATE, gameStore, updateGame } from "../store/game";
  import { toggleModal } from "../store/modal";
  import { navigate } from "svelte-routing";

  let gameState = INITIAL_STATE;

  gameStore.subscribe((store) => {
    gameState = store;

    if (gameState.status !== 'playing') {
      toggleModal({
        title: 'Game ended!',
        description: 'Please refresh the page to join another game!',
        status: 'success',
        handleClose: () => navigate('/', { replace: true })
      });
    }
  });
</script>

<main>
  <h1>Game</h1>
  {#if gameState.id.length}
    <p>Share the code for other player to join!</p>
    <div class="w-96 rounded-lg bg-white">
      <p>{gameState.id}</p>
    </div>
  {/if}

  <Board
    state={gameState.state}
    handleUpdate={updateGame}
  />
</main>

