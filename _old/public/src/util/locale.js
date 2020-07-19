// list of available languages
const languages = [
    "en"
];

// Returns the users locale as two character lowercase ISO string.
export function getLocale() {
    let langs = navigator.languages.slice();

    for(let i = 0; i < langs.length; i++) {
        langs[i] = langs[i].substr(0, 2).toLowerCase(); // e.g.: en-US -> en

        for(let j = 0; j < languages.length; j++) {
            if(langs[i] === languages[j]) {
                return langs[i];
            }
        }
    }

    // use default if none was found
    return "en";
}
