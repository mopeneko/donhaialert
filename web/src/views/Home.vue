<template>
  <div>
    <section class="hero is-dark is-bold">
      <div class="hero-body">
        <div class="container">
          <h1 class="title">ドン廃あらーと</h1>
          <h2 class="subtitle">俺達のためのツール</h2>
        </div>
      </div>
    </section>

    <p class="has-text-centered">さあ、我々の活動を数値化しよう。</p>

    <section>
      <div class="container">
        <b-field>
          <b-input
            v-model="host"
            expanded
            placeholder="インスタンスのホスト名(例: example.com)"
          >
          </b-input>
          <b-button @click="login">ログイン</b-button>
        </b-field>
      </div>
    </section>

    <hr />

    <section class="has-text-centered">
      <p class="title is-4">「ドン廃あらーと」とは</p>
      <p>
        1日のトゥート数・フォロー数・フォロワー数の推移を毎日午前0時に未収載でトゥートするWebアプリです。
      </p>

      <hr />

      <p class="title is-4">開発者</p>
      <p><a href="https//lem0n.cc" target="_blank">もぺねこ</a> です。</p>

      <hr />

      <p class="title is-4">ソースコード</p>
      <p>
        <a href="https://github.com/mopeneko/donhaialert" target="_blank">
          GitHub
        </a>
        にて公開しています。
      </p>
    </section>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class Home extends Vue {
  host = "";

  login() {
    fetch(`https://api.donhaialert.com/auth?host=${this.host}`)
      .then(res => {
        if (res.ok) {
          return res.json();
        }

        alert(`レスポンスが異常です。status code = ${res.status}`);
      })
      .then(j => {
        document.location.href = j["message"];
      })
      .catch(err => {
        alert(err);
      });
  }
}
</script>
