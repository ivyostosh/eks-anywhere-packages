# API Reference

Packages:

- [packages.eks.amazonaws.com/v1alpha1](#packageseksamazonawscomv1alpha1)

# packages.eks.amazonaws.com/v1alpha1

Resource Types:

- [Package](#package)




## Package
<sup><sup>[↩ Parent](#packageseksamazonawscomv1alpha1 )</sup></sup>






Package is the Schema for the package API.

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
      <td>Package</td>
      <td>true</td>
      </tr>
      <tr>
      <td><b><a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.20/#objectmeta-v1-meta">metadata</a></b></td>
      <td>object</td>
      <td>Refer to the Kubernetes API documentation for the fields of the `metadata` field.</td>
      <td>true</td>
      </tr><tr>
        <td><b><a href="#packagespec">spec</a></b></td>
        <td>object</td>
        <td>
          PackageSpec defines the desired state of an package.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#packagestatus">status</a></b></td>
        <td>object</td>
        <td>
          PackageStatus defines the observed state of Package.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Package.spec
<sup><sup>[↩ Parent](#package)</sup></sup>



PackageSpec defines the desired state of an package.

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
        <td><b>packageName</b></td>
        <td>string</td>
        <td>
          PackageName is the name of the package as specified in the bundle.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>config</b></td>
        <td>string</td>
        <td>
          Config for the package.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>packageVersion</b></td>
        <td>string</td>
        <td>
          PackageVersion is a human-friendly version name or sha256 checksum for the package, as specified in the bundle.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>targetNamespace</b></td>
        <td>string</td>
        <td>
          TargetNamespace defines where package resources will be deployed.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Package.status
<sup><sup>[↩ Parent](#package)</sup></sup>



PackageStatus defines the observed state of Package.

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
        <td><b>currentVersion</b></td>
        <td>string</td>
        <td>
          Version currently installed.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b><a href="#packagestatussource">source</a></b></td>
        <td>object</td>
        <td>
          Source associated with the installation.<br/>
        </td>
        <td>true</td>
      </tr><tr>
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
          State of the installation.<br/>
          <br/>
            <i>Enum</i>: initializing, installing, installed, updating, uninstalling, unknown<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b>targetVersion</b></td>
        <td>string</td>
        <td>
          Version to be installed.<br/>
        </td>
        <td>false</td>
      </tr><tr>
        <td><b><a href="#packagestatusupgradesavailableindex">upgradesAvailable</a></b></td>
        <td>[]object</td>
        <td>
          UpgradesAvailable indicates upgraded versions in the bundle.<br/>
        </td>
        <td>false</td>
      </tr></tbody>
</table>


### Package.status.source
<sup><sup>[↩ Parent](#packagestatus)</sup></sup>



Source associated with the installation.

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
        <td><b>digest</b></td>
        <td>string</td>
        <td>
          Digest is a checksum value identifying the version of the package and its contents.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>registry</b></td>
        <td>string</td>
        <td>
          Registry in which the package is found.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>repository</b></td>
        <td>string</td>
        <td>
          Repository within the Registry where the package is found.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          Versions of the package supported.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>


### Package.status.upgradesAvailable[index]
<sup><sup>[↩ Parent](#packagestatus)</sup></sup>



PackageAvailableUpgrade details the package's available upgrades' versions.

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
        <td><b>tag</b></td>
        <td>string</td>
        <td>
          Tag is a specific version number or sha256 checksum for the package upgrade.<br/>
        </td>
        <td>true</td>
      </tr><tr>
        <td><b>version</b></td>
        <td>string</td>
        <td>
          Version is a human-friendly version name for the package upgrade.<br/>
        </td>
        <td>true</td>
      </tr></tbody>
</table>