<template>
    <div id="app">
        <router-view />
    </div>
</template>

<style lang="scss">
@import "./scss/variables.scss";

#app {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: map-get($map: $colors, $key: 2);
}
button {
    background: transparent;
    color: map-get($map: $colors, $key: 4);
    border: none;
    font-size: 16px;
    text-transform: uppercase;
    font-weight: bold;
    padding-top: 24px;
}
body {
    background-color: map-get($map: $colors, $key: 1);
    font-family: Roboto;
    color: map-get($map: $colors, $key: 4);
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
}

li {
    list-style-type: none;
}
button:focus {
    outline: 0px;
}
</style>
