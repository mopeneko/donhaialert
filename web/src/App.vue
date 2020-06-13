<template>
  <div id="app">
    <b-navbar>
      <template slot="brand">
        <b-navbar-item tag="router-link" :to="{ path: '/' }"
          >ドン廃あらーと</b-navbar-item
        >
      </template>

      <template slot="start">
        <b-navbar-item tag="router-link" :to="{ path: '/' }"
          >ホーム</b-navbar-item
        >
      </template>
    </b-navbar>

    <router-view />
  </div>
</template>

<style>
#app {
  font-family: sans-serif;
}
</style>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import "buefy/dist/buefy.css";

@Component
export default class App extends Vue {}
</script>
