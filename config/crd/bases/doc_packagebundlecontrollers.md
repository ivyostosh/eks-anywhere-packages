# API Reference

Packages:

- [packages.eks.amazonaws.com/v1alpha1](#packageseksamazonawscomv1alpha1)

# packages.eks.amazonaws.com/v1alpha1

Resource Types:

- [PackageBundleController](#packagebundlecontroller)




## PackageBundleController
<sup><sup>[↩ Parent](#packageseksamazonawscomv1alpha1 )</sup></sup>






PackageBundleController is the Schema for the packagebundlecontrollers API.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
      <td><b>apiVersion</b></td>
      <td>string</td>
      <td>packages.eks.amazonaws.com/v1alpha1</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b>kind</b></td>
      <td>string</td>
      <td>PackageBundleController</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#packagebundlecontrollerspec">spec</a></b></td>
        <td>object</td>
        <td>
          PackageBundleControllerSpec defines the desired state of PackageBundleController.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#packagebundlecontrollerstatus">status</a></b></td>
        <td>object</td>
        <td>
          PackageBundleControllerStatus defines the observed state of PackageBundleController.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundleController.spec
<sup><sup>[↩ Parent](#packagebundlecontroller)</sup></sup>



PackageBundleControllerSpec defines the desired state of PackageBundleController.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b><a href="#packagebundlecontrollerspecsource">source</a></b></td>
        <td>object</td>
        <td>
          Source of the bundle.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>activeBundle</b></td>
        <td>string</td>
        <td>
          ActiveBundle is name of the bundle from which packages should be sourced.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>logLevel</b></td>
        <td>integer</td>
        <td>
          LogLevel controls the verbosity of logging in the controller.<br/>
          <br/>
            <i>Format</i>: int32<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>privateRegistry</b></td>
        <td>string</td>
        <td>
          PrivateRegistry is the registry being used for all images, charts and bundles<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>upgradeCheckInterval</b></td>
        <td>string</td>
        <td>
          UpgradeCheckInterval is the time between upgrade checks. 
 The format is that of time's ParseDuration.<br/>
          <br/>
            <i>Default</i>: 1d<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>upgradeCheckShortInterval</b></td>
        <td>string</td>
        <td>
          UpgradeCheckShortInterval time between upgrade checks if there is a problem. 
 The format is that of time's ParseDuration.<br/>
          <br/>
            <i>Default</i>: 1h<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### PackageBundleController.spec.source
<sup><sup>[↩ Parent](#packagebundlecontrollerspec)</sup></sup>



Source of the bundle.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>registry</b></td>
        <td>string</td>
        <td>
          Registry portion of an OCI address to the bundle<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Repository portion of an OCI address to the bundle<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### PackageBundleController.status
<sup><sup>[↩ Parent](#packagebundlecontroller)</sup></sup>



PackageBundleControllerStatus defines the observed state of PackageBundleController.

<table>
    <thead>
        <tr>
            <th>Name</th>
            <th>Type</th>
            <th>Description</th>
            <th>Required</th>
        </tr>
    </thead>
    <tbody><tr>
        <td><b>detail</b></td>
        <td>string</td>
        <td>
          Detail of the state.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>state</b></td>
        <td>enum</td>
        <td>
          State of the bundle controller.<br/>
          <br/>
            <i>Enum</i>: ignored, active, disconnected<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>