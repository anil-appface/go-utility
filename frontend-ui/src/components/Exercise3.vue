<template>
  <div>
    Users are stored in MySQL DB. 
    If you are not able to do any operation, Pls check your backend whether it is correctly configured for DB
    <md-table v-model="users" md-card>
      <md-table-toolbar>
        <h1 class="md-title">Users</h1>
        <md-button class="md-primary md-raised" @click="showDialog = true">Create User</md-button>
      </md-table-toolbar>

      <md-table-row slot="md-table-row" slot-scope="{ item }">
        <md-table-cell md-label="Login ID" md-sort-by="loginId" md-numeric>{{ item.loginId }}</md-table-cell>
        <md-table-cell md-label="Full Name" md-sort-by="fullName">{{ item.fullName }}</md-table-cell>
        <md-table-cell md-label="Enabled" md-sort-by="isEnabled">{{ item.isEnabled }}</md-table-cell>
        <md-table-cell md-label="Actions">
            <md-button class="md-icon-button md-raised" @click="editClick(item)">
                <md-icon>edit</md-icon>
             </md-button>
            <md-button class="md-icon-button md-raised md-accent"  @click="deleteUser(item.id)">
                <md-icon>delete</md-icon>
            </md-button>
        </md-table-cell>
      </md-table-row>
    </md-table>

<md-dialog :md-active.sync="showDialog">
      <md-dialog-title>Create User</md-dialog-title>

    <md-field>
      <label>Login ID*</label>
      <md-input v-model="loginId"></md-input>
    </md-field>
    <md-field>
      <label>Full Name*</label>
      <md-input v-model="fullName"></md-input>
    </md-field>
      <md-checkbox v-model="isEnabled">Enabled</md-checkbox>
     
      <md-dialog-actions>
        <md-button class="md-primary" @click="showDialog = false">Close</md-button>
        <md-button class="md-primary" @click="createUser()">Save</md-button>
      </md-dialog-actions>
    </md-dialog>
  </div>

  
</template>

<script>
export default {
  name: "Exercise1",
  data() {
    return {
      users: [],
      showDialog: false,
      isEnabled: true,
      fullName: "",
      loginId: "",
      id:""
    }
  },
  created() {
      this.fetchUsers()
  },
  methods: {
    fetchUsers: function() {
      fetch(`${process.env.VUE_APP_BASE_URL}/users`)
      .then((response) => response.json())
      .then(result => this.users = result);
    },
    deleteUser: function(id) {
      const requestOptions = {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
      };
      fetch(`${process.env.VUE_APP_BASE_URL}/user/${id}`, requestOptions)
      .then(() =>  this.fetchUsers());
    },
    createUser: function() {
      const requestOptions = {
        method: "POST",
        headers: { "Content-Type": "application/json" },
         body: JSON.stringify({ loginId: this.loginId, fullName: this.fullName, id: this.id, isEnabled: this.isEnabled }),
      };
      fetch(`${process.env.VUE_APP_BASE_URL}/user`, requestOptions)
      .then((response) => response.json())
      .then(() => {
          this.fetchUsers()
          this.showDialog = false;
          this.loginId = "";
          this.fullName = "";
          this.isEnabled = true,
          this.id = "";
      });
    },
    editClick(item) {
        this.loginId = item.loginId
        this.isEnabled = item.isEnabled
        this.fullName = item.fullName
        this.id = item.id
        this.showDialog = true
    }
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.md-field, .md-checkbox {
  width: 80%;
  margin: auto;
  margin-top: 10px; 
}

</style>
