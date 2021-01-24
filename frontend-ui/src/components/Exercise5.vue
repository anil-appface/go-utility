<template>
  <div>
      To display all the possible prime numbers upto given input
    <div class="input-divs">
      <md-field>
        <label>Enter number</label>
        <md-input v-model="number"></md-input>
      </md-field>
    </div>
    <md-button class="md-primary" @click="getPrimeNumbers()">Submit</md-button>
    <span>
      {{ result }}
    </span>
  </div>
</template>

<script>
export default {
  name: "Exercise5",
  props: ["number", "result"],
  methods: {
    getPrimeNumbers: function() {
      fetch(`http://localhost:8000/getPrimeNumbers/${this.number}`).then(
        async (response) => {
          this.result = await response.text();
          console.log(this.result);
        }
      );
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.input-divs {
  display: flex;
}

.md-field {
  padding: 10px;
}

.result-div-outside {
  display: flex;
}

.result-div-inside {
  padding: 10px;
  border: 1px solid black;
  margin: 5px;
}
</style>
