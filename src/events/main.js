import {writable, get, derived} from 'svelte/store';

export async function getData(url, method = 'GET') {
    return new Promise((resolve, reject) => {
        fetch(url,{method: method})
            .then((response) => {
                if (!response.ok) {
                    reject(`Error: ${response.status} - ${response.statusText}`);
                }
                return response.json();
            })
            .then((data) => {
                if (data.result) {
                    resolve(data.result);
                } else {
                    resolve(data);
                }
            })
            .catch((error) => {
                reject(`Error: ${error}`);
            });
    });
}

// create me a fetch with a json post body
export async function postData(url, data) {
    return new Promise((resolve, reject) => {
        fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
            .then((response) => {
                if (!response.ok) {
                    reject(`Error: ${response.status} - ${response.statusText}`);
                }
                return response.json();
            })
            .then((data) => {
                if (data.result) {
                    resolve(data.result);
                } else {
                    resolve(data);
                }
            })
            .catch((error) => {
                reject(`Error: ${error}`);
            });
    });
}

// create me a fetch function with a json put body
export async function putData(url, data) {
    return new Promise((resolve, reject) => {
        fetch(url, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        })
            .then((response) => {
                if (!response.ok) {
                    reject(`Error: ${response.status} - ${response.statusText}`);
                }
                return response.json();
            })
            .then((data) => {
                if (data.result) {
                    resolve(data.result);
                } else {
                    resolve(data);
                }
            })
            .catch((error) => {
                reject(`Error: ${error}`);
            });
    });
}

function createUserEvents() {
    const {subscribe, set} = writable({isLoggedIn: false, userData : {}});

    return {
        subscribe,
        loadUser: async (callback) => {
            await fetch(`/api/auth/user-info`)
                .then(response => {
                    if (!response.ok) {
                        throw new Error('Failed to fetch user info during event pipeline');
                    }
                    return response.json();
                })
                .then(data => {
                    set({isLoggedIn: true, userData: data.result});
                    if (typeof callback === "function") {
                        callback(data.result);
                    }
                })
                .catch(error => {
                    set({isLoggedIn: false, userData: {}});
                });
        },
    };
}

export const userEvents = createUserEvents();

// Create teh growl events items
function createGrowlEvents() {
    const {subscribe, set} = writable({isShown: false, message: "No Content", type: "info"});

    return {
        subscribe,
        setContent: (inContent, inType) => set({
            message: inContent,
            isShown: true,
            type: inType
        }),
        hide: () => set({message: "No Content", isShown: false, type: ""}),
    };
}

export const growlEvents = createGrowlEvents();