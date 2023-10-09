/**
 * This function will redirect the user to a local route inside the website.
 * 
 * @param location the path to the route.
 * @returns void
 */
export const redirectLocal = (location: string) => () => {
    window.location.href = "/" + location;
}

/**
 * This function will redirect the user to an external website.
 * 
 * @param location the path to the route.
 * @returns void
 */
export const redirectExternal = (location: string) => () => {
    window.location.href = "https://" + location;
}