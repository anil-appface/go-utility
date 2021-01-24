<template>
  <div class="page-container">

     <md-tabs>
      <md-tab id="tab-exercise1" md-label="Exercise 1">
        <Exercise1 />
      </md-tab>

      <md-tab id="tab-exercise2" md-label="Exercise 2">
        <Exercise2 />  
      </md-tab>

      <md-tab id="tab-exercise3" md-label="Exercise 3" >
         <Exercise3 />  
     </md-tab>

      <md-tab id="tab-exercise4" md-label="Exercise 4">
       <Exercise4 />  
       </md-tab>

       <md-tab id="tab-exercise5" md-label="Exercise 5">
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
import Exercise3 from './components/Exercise3'
import Exercise4 from './components/Exercise4'
import Exercise5 from './components/Exercise5'

export default {
  name: 'Waterfall',
  components: {
    Exercise1,
    Exercise2,
    Exercise3,
    Exercise4,
    Exercise5
  }
}

</script>
