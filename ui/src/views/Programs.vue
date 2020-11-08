<template>
  <v-container>
    <v-row>
      <v-col sm="1" offset-lg="1">
        <Sidebar/>
      </v-col>
      <v-col sm="9" lg="9">
        <v-data-table
            :headers="headers"
            :items="programs"
            sort-by="date"
            class="elevation-1"
        >
          <template v-slot:top>
            <v-toolbar
                flat
            >
              <v-toolbar-title>CRUD</v-toolbar-title>
              <v-divider
                  class="mx-4"
                  inset
                  vertical
              ></v-divider>
              <v-spacer></v-spacer>
              <v-dialog
                  v-model="dialog"
                  fullscreen
                  hide-overlay
                  transition="dialog-bottom-transition"
              >
                <template v-slot:activator="{ on, attrs }">
                  <v-btn
                      color="primary"
                      dark
                      class="mb-2"
                      v-bind="attrs"
                      v-on="on"
                  >
                    New Item
                  </v-btn>
                </template>
                <v-card>
                  <v-card-title>
                    <span class="headline">{{ formTitle }}</span>
                  </v-card-title>

                  <v-card-text>
                    <v-container>
                      <v-row>
                        <v-col
                            cols="12"
                            sm="3"
                            md="3"
                        >
                          <v-text-field
                              v-model="editedItem.date"
                              label="Date"
                              readonly
                          ></v-text-field>
                        </v-col>
                        <v-col
                            cols="12"
                            sm="3"
                            md="3"
                        >
                          <v-text-field
                              v-model="editedItem.title"
                              label="Title"
                          ></v-text-field>
                        </v-col>
                        <v-col
                            cols="12"
                            sm="3"
                            md="3"
                        >
                          <v-text-field
                              v-model="editedItem.tags"
                              label="Tags"
                          ></v-text-field>
                        </v-col>
                        <v-col
                            cols="12"
                            sm="6"
                            md="6"
                        >
                          <v-textarea
                              v-model="editedItem.text"
                              label="Text"
                              auto-grow
                          ></v-textarea>
                        </v-col>

                      </v-row>
                    </v-container>
                  </v-card-text>

                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                        color="blue darken-1"
                        text
                        @click="close"
                    >
                      Cancel
                    </v-btn>
                    <v-btn
                        color="blue darken-1"
                        text
                        @click="save"
                    >
                      Save
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
              <v-dialog v-model="dialogDelete" max-width="500px">
                <v-card>
                  <v-card-title class="headline">Are you sure you want to delete this item?</v-card-title>
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="closeDelete">Cancel</v-btn>
                    <v-btn color="blue darken-1" text @click="deleteItemConfirm">OK</v-btn>
                    <v-spacer></v-spacer>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            </v-toolbar>
          </template>
          <template v-slot:item.actions="{ item }">
            <v-icon
                small
                class="mr-2"
                @click="editItem(item)"
            >
              mdi-pencil
            </v-icon>
            <v-icon
                small
                @click="deleteItem(item)"
            >
              mdi-delete
            </v-icon>
          </template>
          <template v-slot:no-data>
            <v-btn
                color="primary"
                @click="initialize"
            >
              Reset
            </v-btn>
          </template>
        </v-data-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Sidebar from "@/components/programs/Sidebar"
import programs from "@/store/programs";

export default {
  name: "Programs",
  components: {Sidebar},
  data: () => ({
    dialog: false,
    dialogDelete: false,
    headers: [
      {text: 'Date', align: 'start', value: 'date'},
      {text: 'Text', value: 'text'},
      {text: 'Title', value: 'title',},
      {text: 'Tags', value: 'tags'},
      {text: 'Actions', value: 'actions', sortable: false},
    ],
    programs: [],
    editedIndex: -1,
    editedItem: {
      title: '',
      text: '',
      tags: '',
      date: null,
    },
    defaultItem: {
      title: '',
      text: '',
      tags: '',
      date: null,
    },
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
    }
  },

  watch: {
    dialog(val) {
      val || this.close()
    },
    dialogDelete(val) {
      val || this.closeDelete()
    },
  },

  created() {
    this.initialize()
  },

  methods: {
    initialize() {
      this.programs = programs.map(item => {
        return {
          ...item,
          title: item.text.substr(0, 30),
        }
      })
    },

    editItem(item) {
      this.editedIndex = this.programs.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    deleteItem(item) {
      this.editedIndex = this.programs.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialogDelete = true
    },

    deleteItemConfirm() {
      this.programs.splice(this.editedIndex, 1)
      this.closeDelete()
    },

    close() {
      this.dialog = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    closeDelete() {
      this.dialogDelete = false
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      })
    },

    save() {
      if (this.editedIndex > -1) {
        Object.assign(this.programs[this.editedIndex], this.editedItem)
      } else {
        this.programs.push(this.editedItem)
      }
      this.close()
    },
  },
}
</script>
