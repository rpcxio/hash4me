<template>
  <div id="app" class="container">
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>挖矿算力演示</h1>

        <form v-on:submit.prevent="calc">
          <div class="form-group">
            <input v-model="targetValue" type="text" id="target-input" placeholder="输入一个1到10的难度" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">计算!</button>
          </div>
        </form>

         <p v-text = 'hash'></p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'App',

  data() { return {
    hash: '',
  } },

  methods: {
    calc() {
      axios.post("http://localhost:9981/api/hash", {
        target: this.targetValue,
      })
              .then((response) => {
                this.hash = response.data.hash;
              })
              .catch((error) => {
                window.alert(`The API returned an error: ${error}`);
              })
    }
  }
}
</script>
