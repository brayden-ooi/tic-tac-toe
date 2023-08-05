<script lang="ts">
  import type { gameType } from '../store/game';

  export let state: gameType["state"] = Array(3).fill(Array(3).fill(''));
  export let handleUpdate = (x: number, y: number) => {};

</script>

<div class="board-container">
  {#each state as row, y}
    {#each row as col, x}
      <div
        class="board-cell"
        on:click={() => handleUpdate(x, y)}>
        <span>
          {#switch col}
            {:case "X"}
              ❌
            {:case "O"}
              🟢
            {:default}
          {/switch}
        </span>
      </div>
    {/each}
  {/each}
</div>

<style>
  .board-container {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr;
    grid-template-rows: 1fr 1fr 1fr;
  }

  .board-cell {
    height: 100px;
    width: 100px;
    background-color: white;
    border: 1px black solid;
    padding: 50px;
    padding-top: 55px;
    padding-bottom: 45px;
  }

  .board-cell span {
    color: black;
    font-size: 50px;
  }
</style>