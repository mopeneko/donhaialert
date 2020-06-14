<template>
  <div class="has-text-centered">
    <b-button @click="changeSettings('POST')">登録</b-button>
    <b-button @click="changeSettings('DELETE')">解除</b-button>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class Settings extends Vue {
  changeSettings(method: string) {
    fetch("https://api.donhaialert.com/settings", {
      method
    })
      .then(res => {
        if (res.ok) {
          this.$router.push("/");
        }
        throw new Error("");
      })
      .catch(() => {
        this.$router.push("/error");
      });
  }
}
</script>
