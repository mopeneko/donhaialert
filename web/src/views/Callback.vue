<template>
  <div></div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class Callback extends Vue {
  beforeCreate() {
    fetch("https://api.donhaialert.com/callback", {
      method: "POST",
      headers: {
        "Content-Type": "application/json; charset=utf-8"
      },
      body: JSON.stringify({
        code: this.$route.query.code,
        state: this.$route.query.state
      })
    })
      .then(res => {
        if (res.ok) {
          return res.json();
        }

        throw new Error("");
      })
      .then(() => {
        this.$router.push("settings");
      })
      .catch(() => {
        this.$router.push("error");
      });
  }
}
</script>
