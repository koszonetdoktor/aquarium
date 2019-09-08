<template>
    <div id="app">
        <router-view />
    </div>
</template>

<style>
#app {
    font-family: "Avenir", Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
}
button {
    background: transparent;
    color: #93b7cc;
    border: none;
    font-size: 16px;
    text-transform: uppercase;
    font-weight: bold;
    padding-top: 24px;
}
body {
    background-color: #132434;
    font-family: Roboto;
    color: #93b7cc;
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
