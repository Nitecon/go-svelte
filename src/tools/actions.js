import { derived } from 'svelte/store';

export function contentTransition(statusStore) {
    const transitionKey = derived(statusStore, $status => {
        if ($status === "" || $status === null) {
            return 'form';
        } else if ($status === 'success') {
            return 'success';
        } else {
            return 'error';
        }
    });

    return {
        fadeTransition: {
            delay: 250,
            duration: 250
        },
        transitionKey
    };
}

export function goBack() {
    window.history.back();
}

export function reloadMe() {
    window.location.reload();
}

export async function copyToClipboard(id, type) {
    const fullURL = window.location.origin + "/api/" + type + "/" + id + "/data";
    await navigator.clipboard.writeText(fullURL);
}

export async function copyTextToClipboard(data) {
    await navigator.clipboard.writeText(data);
}

