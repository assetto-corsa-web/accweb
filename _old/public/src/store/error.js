export const ErrorStore = {
    state: {
        errorMsg: null
    },
    mutations: {
        toast(state, error) {
            state.errorMsg = error;
            state.toastTimeout = setTimeout(() => {
                state.errorMsg = null;
            }, 3000);
        }
    }
};
