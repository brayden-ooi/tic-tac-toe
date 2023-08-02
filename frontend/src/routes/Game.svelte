<script lang="ts">
  import Board from "../lib/Board.svelte";
  import { INITIAL_STATE, gameStore, updateGame } from "../store/game";

  let gameState = INITIAL_STATE;

  gameStore.subscribe((store) => {
    gameState = store;
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

