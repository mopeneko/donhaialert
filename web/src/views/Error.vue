<template>
  <div class="has-text-centered">
    <p>エラーが起きました。</p>
    <b-button tag="router-link" to="/" type="link">ホームへ戻る</b-button>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class Error extends Vue {}
</script>
