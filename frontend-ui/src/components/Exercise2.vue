<template>
  <div>
      <div class="input-divs">
    <md-field>
      <label>Column Start</label>
      <md-input v-model="columnStart"></md-input>
    </md-field>
    <md-field>
      <label>Number of Rows</label>
      <md-input v-model="rowCount"></md-input>
    </md-field>
    <md-field>
      <label>Number of Columns</label>
      <md-input v-model="colCount"></md-input>
    </md-field>
    </div>
    <md-button class="md-primary" @click="calculate()">Submit</md-button>

    <div>
         <div class="result-div-outside" v-for="(v,i) in result" :key="i" >
            <div class="result-div-inside" v-for="(ch,j) in v" :key="j">
                {{ch}}
             </div>
        </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Exercise1',
  props: ['columnStart', 'rowCount', 'colCount', 'result'],
  methods: {
    calculate: function() {

        console.log(this.columnStart, this.rowCount, this.colCount)
        if (!this.columnStart || !this.rowCount || !this.colCount) {
            alert('error')
            return
        }

        let colStart =  this.columnStart.toUpperCase()
        console.log(colStart)
        let prev = this.columnStart
        this.result = []
        for (let index = 0; index < this.rowCount; index++) {
            let arr = []
            for (let j = 0; j < this.colCount; j++) {
                arr.push(prev)
                prev = this.nextString(prev);
            }
            this.result.push(arr)
        }
        console.log(this.result)
    }, 
    nextString: function(key) {
         if (key === 'Z' || key === 'z') {
            return String.fromCharCode(key.charCodeAt() - 25) + String.fromCharCode(key.charCodeAt() - 25); // AA or aa
        } else {
            var lastChar = key.slice(-1);
            var sub = key.slice(0, -1);
            if (lastChar === 'Z' || lastChar === 'z') {
                // If a string of length > 1 ends in Z/z,
                // increment the string (excluding the last Z/z) recursively,
                // and append A/a (depending on casing) to it
                 return this.nextString(sub) + String.fromCharCode(lastChar.charCodeAt() - 25);
            } else {
            // (take till last char) append with (increment last char)
                return sub + String.fromCharCode(lastChar.charCodeAt() + 1);
            }
        }
    }
  }
}
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
