import { writable } from 'svelte/store';

  function createUserStore() {
	const { subscribe, set, update } = writable({
        name: "marcus",
        loggedIn: false,
      });

	return {
		subscribe,
		logIn: () => update(user => {
            user.loggedIn = true;            
            return user;
        }),
		logOut: () => update(user => {
            user.loggedIn = false;
            return user;
        })
	};
}

export const user = createUserStore();