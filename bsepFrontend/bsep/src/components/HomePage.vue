<template>
  <div>
    <div class="row">
      <div class="col-12">
        <h1 v-if="authority !== 'admin'">My certificates</h1>
        <h1 v-if="authority === 'admin'">All certificates</h1>
      </div>
      <div class="col d-flex justify-content-end" style="margin-right: 10%">
        <vs-tooltip>
          <vs-button @click="openDialog()" v-if="authority !== 'endEntity'" :disabled="selectedCert === null && authority !== 'admin'" primary size="xl" icon>
            Issue certificate
          </vs-button>
          <template #tooltip>
            <p v-if="authority !== 'admin'">First select certificate to sign with</p>
            <p v-if="authority === 'admin'">First select certificate to sign with or issue self-signed certificate</p>
          </template>
        </vs-tooltip>
      </div>
    </div>

    <vs-dialog @close="resetForm()" auto-width v-model="dialog">
      <template #header>
        <h4 v-if="selectedCert" class="not-margin">
          New certificate
        </h4>
        <h4 v-if="!selectedCert" class="not-margin me-3 ms-3">
          New self-signed certificate
        </h4>
      </template>
      <div class="con-form" style="display: inline-block">
        <div class="row">
          <div class="col">
            <vs-input required class="mt-2" primary v-model="commonNameSubject" label-placeholder="Subject common name"/>
            <vs-input required class="mt-4" primary v-model="nameSubject" label-placeholder="Subject first name"/>
            <vs-input required class="mt-4" primary v-model="surnameSubject" label-placeholder="Subject last name"/>
            <vs-input required class="mt-4" primary v-model="usernameSubject" label-placeholder="Subject username"/>
            <vs-input required class="mt-4" primary v-model="countrySubject" label-placeholder="Subject country"/>
          </div>
          <div class="col">
            <vs-select required class="mt-2" v-if="selectedCert !== null"
                label-placeholder="Select basic constraint"
                v-model="basicConstraints">
              <vs-option selected label="CA" value="ca">
                CA
              </vs-option>
              <vs-option label="End entity" value="endEntity">
                End entity
              </vs-option>
            </vs-select>
            <vs-input required class="mt-4" :min="getMinDate()" type="date" primary v-model="validFrom" label="Valid from"/>
            <vs-input required class="mt-4" :disabled="validFrom === ''" primary type="date" :min="getMinDateValidTo()" :max="getMaxDate()" v-model="validTo" label="Valid to"/>
            <vs-select required class="mt-4"
                filter
                multiple
                collapse-chips
                label="Key usage"
                v-model="keyUsages">
              <vs-option label="Encipher Only" value="1">
                Encipher Only
              </vs-option>
              <vs-option label="CRL Signing" value="2">
                CRL Signing
              </vs-option>
              <vs-option label="Key Certificate Signing" value="4">
                Key Certificate Signing
              </vs-option>
              <vs-option label="Key Agreement" value="8">
                Key Agreement
              </vs-option>
              <vs-option label="Data Encipherment" value="16">
                Data Encipherment
              </vs-option>
              <vs-option label="Key Encipherment" value="32">
                Key Encipherment
              </vs-option>
              <vs-option label="Non-repudiation" value="64">
                Non-repudiation
              </vs-option>
              <vs-option label="Digital Signature" value="128">
                Digital Signature
              </vs-option>
              <vs-option label="Decipher Only" value="32768">
                Decipher Only
              </vs-option>
            </vs-select>
          </div>
        </div>
      </div>
      <template #footer>
        <div class="footer-dialog">
          <vs-button :disabled="isFormInvalid()" @click="issueCertificate()" block>
            Issue certificate
          </vs-button>
        </div>
      </template>
    </vs-dialog>
    <div class="row mt-3">
      <div v-for="c in myCertificates" :key="c.serialNumberSubject" class="col d-flex justify-content-center">
        <vs-card :id="c.serialNumberSubject" @click="setSelectedCert(c)">
          <template #title>
            <h3>{{c.commonNameSubject}}</h3>
          </template>
          <template #img>
            <img :src="'/' + c.authoritySubject + '.svg'" alt="" width="200" height="200">
          </template>
          <template #text>
            <p style="text-align: left">
              <strong>Subject username:</strong> {{c.usernameSubject}}<br>
              <strong>Subject name:</strong> {{c.nameSubject}}<br>
              <strong>Subject surname:</strong> {{c.surnameSubject}}<br>
              <strong>Subject country:</strong> {{c.countrySubject}}<br>
              <strong>Subject serial number:</strong> {{c.serialNumberSubject}}<br>
              <strong>Subject authority:</strong> {{c.authoritySubject}}<br>
              <strong>Key usages:</strong> {{keyUsagePipe(c.keyUsages)}}<br>


              <strong>Valid from:</strong> {{c.startDate | date}}<br>
              <strong>Valid until:</strong> {{c.endDate | date}}<br>

              <strong>Issuer username:</strong> {{c.usernameIssuer}}<br>
              <strong>Issuer common name:</strong> {{c.commonNameIssuer}}<br>
            </p>
          </template>
          <template #interactions>
            <vs-button @click="revokeCertificate(c)" v-if="authority === 'admin'" danger icon>
              Revoke
            </vs-button>
            <vs-button @click="downloadCertificate(c)" class="btn-primary" primary>
              Download
            </vs-button>
          </template>
        </vs-card>
      </div>
    </div>
    <div v-if="authority !== 'endEntity' && authority !== 'admin'" class="row mt-5"><h1>Issued certificates</h1></div>
    <div v-if="authority !== 'endEntity' && authority !== 'admin'" class="row mb-5 mt-5">
      <div v-for="c in issued" :key="c.serialNumberSubject" class="col d-flex justify-content-center">
      <vs-card>
        <template #title>
          <h3>{{c.commonNameSubject}}</h3>
        </template>
        <template #img>
          <img :src="'/' + c.authoritySubject + '.svg'" alt="" width="200" height="200">
        </template>
        <template #text>
          <p style="text-align: left">
            <strong>Subject username:</strong> {{c.usernameSubject}}<br>
            <strong>Subject name:</strong> {{c.nameSubject}}<br>
            <strong>Subject surname:</strong> {{c.surnameSubject}}<br>
            <strong>Subject country:</strong> {{c.countrySubject}}<br>
            <strong>Subject serial number:</strong> {{c.serialNumberSubject}}<br>
            <strong>Subject authority:</strong> {{c.authoritySubject}}<br>
            <strong>Key usages:</strong> {{keyUsagePipe(c.keyUsages)}}<br>


            <strong>Valid from:</strong> {{c.startDate | date }}<br>
            <strong>Valid until:</strong> {{c.endDate | date }}<br>

            <strong>Issuer username:</strong> {{c.usernameIssuer}}<br>
            <strong>Issuer common name:</strong> {{c.commonNameIssuer}}<br>
          </p>
        </template>
      </vs-card>
    </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment"

export default {
  name: "HomePage",
  data() {
    return {
      authority: "endEntity",
      user: null,
      myCertificates: [],
      issued: [],
      selectedCert: null,
      path: "",
      dialog: false,
      commonNameSubject: "",
      nameSubject: "",
      surnameSubject: "",
      countrySubject: "",
      usernameSubject: "",
      basicConstraints: 'ca',
      validFrom: '',
      validTo: '',
      keyUsages: [],
      keyUsageMap: {},
    };
  },
  async mounted() {
    this.keyUsageMap = {
      128: 'digitalSignature',
      64: 'nonRepudiation',
      32: 'keyEncipherment',
      16: 'dataEncipherment',
      8: 'keyAgreement',
      4: 'keyCertSign',
      2: 'cRLSign',
      1: 'encipherOnly',
      32768: 'decipherOnly',
    };
    this.user = this.$store.getters.user;
    this.authority = this.$store.getters.user?.role?.authority;
    if(this.authority === 'admin') {
      const response = await axios.get(`${process.env.VUE_APP_BACKEND}/admin/getAllCertificates`);
      if(response.data) {
        this.myCertificates = response.data;
      }
      return;
    }
    const response = await axios.get(`${process.env.VUE_APP_BACKEND}/user/${this.user?.username}/certificate`);
    if(response.data) {
      console.table(response.data);
      for(let i = 0; i < response.data.length; i++) {
        if(response.data[i].usernameSubject !== this.user?.username) {
          this.issued.push(response.data[i]);
        } else {
          this.myCertificates.push(response.data[i]);
          const authorities = this.myCertificates.flatMap(a => a.authoritySubject);
          if(authorities.indexOf('root') > -1) this.authority = 'root';
          else if(authorities.indexOf('ca') > -1) this.authority = 'ca';
        }
      }
    }
  },
  methods: {
    resetForm(){
      this.commonNameSubject = "";
      this.nameSubject = "";
      this.surnameSubject = "";
      this.countrySubject = "";
      this.usernameSubject = "";
      if(this.authority === 'admin')
        this.basicConstraints = 'root';
      else this.basicConstraints = 'ca';
      this.validFrom = '';
      this.validTo = '';
      this.keyUsages = [];
    },

    openDialog() {
      this.dialog = true;
      if(this.selectedCert === null)
        this.basicConstraints = 'root';
      else this.basicConstraints = 'ca';
    },

    isFormInvalid(){
      return this.commonNameSubject === "" || this.nameSubject === "" || this.surnameSubject === "" || this.countrySubject === "" || this.usernameSubject === "" || this.basicConstraints === "" || this.validFrom === "" || this.validTo === "";
    },

    async issueCertificate() {
     if(this.isFormInvalid()) {
        this.$vs.notification({
          title: 'Error',
          text: 'Please fill all the fields',
          color: 'danger',
          position: 'top-right'
        });
        return;
      }
     if(this.selectedCert === null) this.basicConstraints = 'root';
      const certificate = {
        commonNameSubject: this.commonNameSubject,
        nameSubject: this.nameSubject,
        surnameSubject: this.surnameSubject,
        countrySubject: this.countrySubject,
        usernameSubject: this.usernameSubject,
        authoritySubject: this.basicConstraints,
        startDate: this.validFrom,
        endDate: this.validTo,
        keyUsages: this.keyUsages,
      };
      if(this.selectedCert === null){
        certificate.commonNameIssuer = this.commonNameSubject;
        certificate.nameIssuer = this.nameSubject;
        certificate.surnameIssuer = this.surnameSubject;
        certificate.countryIssuer = this.countrySubject;
        certificate.usernameIssuer = this.usernameSubject;
        certificate.authorityIssuer = this.basicConstraints;
      } else {
        certificate.commonNameIssuer = this.selectedCert.commonNameSubject;
        certificate.nameIssuer = this.selectedCert.nameSubject;
        certificate.surnameIssuer = this.selectedCert.surnameSubject;
        certificate.countryIssuer = this.selectedCert.countrySubject;
        certificate.usernameIssuer = this.selectedCert.usernameSubject;
        certificate.authorityIssuer = this.selectedCert.authoritySubject;
        certificate.serialNumberIssuer = this.selectedCert.serialNumberSubject;
      }

      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/admin/createCertificate`, certificate);
      if(response.data) {
        this.resetForm();
        this.dialog = false;
        if(this.authority === 'admin')
          this.myCertificates.push(response.data);
        else
          this.issued.push(response.data);
        this.$vs.notification({
          title: 'Success',
          text: 'Certificate issued successfully!',
          color: 'success',
          position: 'top-right'
        });
      }
      else {
        this.$vs.notification({
          title: 'Error',
          text: 'Error while issuing certificate!',
          color: 'danger',
          position: 'top-right'
        });
      }
    },
    async downloadCertificate(c){
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/admin/downloadCertificate`, {
        serialNumberSubject: c.serialNumberSubject,
        authoritySubject: c.authoritySubject,
      });
      if(response.data){
        var a = document.createElement('a');
        var blob = new Blob([response.data], {'type': 'application/octet-stream'});
        a.href = window.URL.createObjectURL(blob);
        a.download = `${c.serialNumberSubject}.crt`;
        a.click();
      }
    },
    setSelectedCert(c){
      if(c.authoritySubject === 'endEntity') return;
      if(this.selectedCert && c.serialNumberSubject !== this.selectedCert.serialNumberSubject)
        document.getElementById(this.selectedCert.serialNumberSubject).firstElementChild.style.backgroundColor = 'white';
      const card = document.getElementById(c.serialNumberSubject)?.firstElementChild
      if(this.selectedCert?.serialNumberSubject === c.serialNumberSubject) {
        card.style.backgroundColor = 'white'
        this.selectedCert = null;
        return
      }
      this.selectedCert = c;
      card.style.backgroundColor = 'cadetblue'
    },

    getMinDate(){
      if(moment(this.selectedCert?.startDate).isAfter(moment()))
        return moment(this.selectedCert?.startDate).add(1, 'days').format('YYYY-MM-DD');
      return moment().format('YYYY-MM-DD');
    },

    getMaxDate(){
      if(this.selectedCert) {
        return moment(this.selectedCert.endDate).subtract('1', 'days').format('YYYY-MM-DD');
      }
      else return null;
    },

    getMinDateValidTo(){
      if(this.validFrom !== '') {
        return moment(this.validFrom).add('1', 'days').format('YYYY-MM-DD');
      }
      else return moment().format('YYYY-MM-DD');
    },

    keyUsagePipe(value){
      return value.map(u => this.keyUsageMap[u]).join(', ');
    },
    async revokeCertificate(c){
      const response = await axios.post(`${process.env.VUE_APP_BACKEND}/admin/revokeCertificate`, {
        serialNumberSubject: c.serialNumberSubject,
        authoritySubject: c.authoritySubject,
      });
      if(response.data){
        this.$vs.notification({
          title: 'Success',
          text: 'Certificate revoked successfully!',
          color: 'success',
          position: 'top-right'
        });
        if(this.authority === 'admin') {
          const response = await axios.get(`${process.env.VUE_APP_BACKEND}/admin/getAllCertificates`);
          if(response.data) {
            this.myCertificates = response.data;
          }
          return;
        }
        const response = await axios.get(`${process.env.VUE_APP_BACKEND}/user/${this.user?.username}/certificate`);
        if(response.data) {
          for(let i = 0; i < response.data.length; i++) {
            if(response.data[i].usernameSubject !== this.$route.params.user) {
              this.issued.push(response.data[i]);
            } else {
              this.myCertificates.push(response.data[i]);
            }
          }
        }
      }
    },
  },

  filters: {
    date(value) {
      return moment(value).format('MM/DD/YYYY');
    },

  },
  props: {

  }
}
</script>

<style>
.vs-select__label--hidden.vs-select__label--placeholder{
  margin-bottom: 5%;
}
.vs-select__label--label{
  margin-bottom: 5%;
}
</style>