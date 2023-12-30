import type { Mode } from './types/mode';
import { writable } from 'svelte/store';

export const mode = writable<Mode>('light');

export const toggleMode = () => {
	mode.update((m) => (m === 'light' ? 'dark' : 'light'));
};
