<template>
  <div class="page-container">

     <md-tabs>
      <md-tab id="tab-home" md-label="Exercise 1">
        <Exercise1 />
      </md-tab>

      <md-tab id="tab-pages" md-label="Exercise 2">
        <Exercise2 />  
      </md-tab>

      <md-tab id="tab-posts" md-label="Posts" >
        Posts tab
        <p>Qui, voluptas repellat impedit ducimus earum at ad architecto consectetur perferendis aspernatur iste amet ex tempora animi, illum tenetur quae assumenda iusto.</p>
      </md-tab>

      <md-tab id="tab-favorites" md-label="Exercise 4">
       <Exercise4 />  
       </md-tab>

       <md-tab id="tab" md-label="Exercise 5">
       <Exercise5 />  
       </md-tab>
    </md-tabs>

  </div>
</template>

<style lang="scss" scoped>

  @import url("https://fonts.googleapis.com/css?family=Material+Icons");

  .page-container {
    max-width: 800px !important;
    margin: auto;
  }
</style>

<script>

import Exercise1 from './components/Exercise1'
import Exercise2 from './components/Exercise2'
import Exercise4 from './components/Exercise4'
import Exercise5 from './components/Exercise5'

export default {
  name: 'Waterfall',
  components: {
    Exercise1,
    Exercise2,
    Exercise4,
    Exercise5
  }
}

</script>
