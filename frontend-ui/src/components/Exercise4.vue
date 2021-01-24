<template>
  <div>
    <div class="input-divs">
      <md-field>
        <label>Enter Date in format '2020-02-18T01:50'</label>
        <md-input v-model="date"></md-input>
      </md-field>
    </div>
    <md-button class="md-primary" @click="getLastDayofMonth()">Submit</md-button>
    <span>
      {{ result }}
    </span>
  </div>
</template>

<script>
export default {
  data(){
    return{
      date: undefined,
      result: undefined
    }
  },
  methods: {
    getLastDayofMonth: function() {
      const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ Date: this.date }),
      };
      fetch(`${process.env.VUE_APP_BASE_URL}/getLastDayofMonth`, requestOptions).then(
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
