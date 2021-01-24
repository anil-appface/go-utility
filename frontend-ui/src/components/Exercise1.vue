<template>
  <div>
    <md-field>
      <label>Enter URL for web scrape</label>
      <md-input v-model="urlValue"></md-input>
    </md-field>
    <md-button class="md-primary" @click="getAllWordsWithCount()">Submit</md-button>

    <div>
    
    <div class="result-div" v-for="(v,i) in result" :key="i" >
         {{i}} - {{v}}
    </div>

    
    </div>
  </div>
</template>

<script>
export default {
  name: 'Exercise1',
  data() {
    return {
      result: {},
      urlValue:""
    }
  },
  methods: {
    getAllWordsWithCount: function() {
      //get all 
      console.log(this.urlValue)
       fetch(`http://localhost:8000/webscrape?url=${this.urlValue}`)
    .then(response => response.json())
    .then(data => (
      this.result = data
    ));
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.result-div {
    border: 1px solid black;
    padding: 10px;
    display: inline-flex;
    margin: 4px;
}
</style>
