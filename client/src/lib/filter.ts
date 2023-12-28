import type { FilterOption } from '$lib/types/filter';
import { writable } from 'svelte/store';

export const filter = writable<FilterOption>('all');