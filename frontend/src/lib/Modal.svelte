<script lang="ts">
  import { INITIAL_STATE, store, toggleModalOff } from '../store';

  // modal state
  let modalState = INITIAL_STATE.modal;

  store.subscribe((store) => {
    modalState = store.modal;
  });

  let dialogElement: HTMLDialogElement;
  
  $: if (modalState.isOpen) {
    dialogElement?.showModal();
  } else {
    dialogElement?.close();
  }
</script>

<dialog bind:this={dialogElement}>
  <p>{modalState.title}</p>
  <p>{modalState.description}</p>
  <button on:click={toggleModalOff}>Close</button>
</dialog>