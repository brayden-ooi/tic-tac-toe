import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'
import switchCase from 'svelte-switch-case';

export default {
  // Consult https://svelte.dev/docs#compile-time-svelte-preprocess
  // for more information about preprocessors
  preprocess: [
    vitePreprocess(),
    switchCase(),
  ],
}
