<script lang="ts">
  import type { stateType } from '../store';

  export let state: stateType["game"]["state"] = Array(3).fill(Array(3).fill(''));

</script>

<div class="board-container">
  {#each state as row}
    {#each row as col}
      <div class="board-cell">
        <span>{col}</span>
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
    height: 200px;
    width: 200px;
    background-color: white;
    border: 1px black solid;
  }

  .board-cell span {
    color: black;
  }
</style>