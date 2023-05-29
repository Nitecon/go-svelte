export function ucfirst(str) {
    return str.charAt(0).toUpperCase() + str.slice(1);
}
export default function getUUIDFromUrl(){
    const hash = window.location.hash.replace(/^#/, '');
    const uuid = hash.split('/')[1];
    return uuid;
}