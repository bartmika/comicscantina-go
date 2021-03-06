Comics Cantina API Reference
======

## Developers Notes

1. To help make the next few API endpoints easy to type, save your token to the console.

    ```bash
    COMICS_WS_API_TOKEN='YOUR_TOKEN'
    ```

2. You will notice ``http`` used in the sample calls through this document, this is the ``Python`` command line application called ``HTTPie``. Download the command line application by following [these instructions](https://httpie.org/).


3. If you are going to make any contributions, please make sure your edits follow the [API documentation standard](https://gist.github.com/iros/3426278) for this document; in addition, please read [Googles API Design Guide](https://cloud.google.com/apis/design/) for further consideration.


## Get API Version
Returns the version information of Comics Cantina. This is a useful endpoint to call when you are setting up your project and you want to confirm you are able to communicate with the web-service.


* **URL**

  ``/api/v1/public/version``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "Service": "v0.1",
        "API: 'v1"
    }
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/public/version
  ```


## Register
Submit registration details into our system to automatically create a *user* account. System return the *user* details and authentication *token*.

Created *user* accounts are automatically granted access to the system even though these accounts have not had their email verified. The system sends a verification email after creation and if the *user* does not verify in the allotted timespan, their account gets locked.

It's important to note that emails must be unique and passwords strong or else validation errors get returned.

* **URL**

  ``/api/v1/public/register``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * email
  * password
  * first_name
  * last_name


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkyOTY1MDAsInVzZXJfaWQiOjF9.QN9dyWL2dlxKgkm0xbQAmnaI6_4amHcSfqUGQ6pZbxM",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Email is not unique.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/public/register \
    email=bart@mikasoftware.com \
    password=YOUR_PASSWORD \
    first_name=Bart \
    last_name=Mika
  ```


## Login
Returns the *user profile* and authentication *token* upon successful login in.

* **URL**

  ``/api/v1/public/login``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * email
  * password


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkyOTg1MDYsInVzZXJfaWQiOjF9.HrwHvfL4-1pMe7EcXEzlsxciFgK0xf2uC8BV1kfLT_c",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Email or password is incorrect.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/public/login \
    email=bart@mikasoftware.com \
    password=YOUR_PASSWORD
  ```


## List Public Organizations
Returns paginated list of all the *organizations* which have been approved by the staff of [**Lucha Comics** ](https://luchacomics.com/) for public viewing. Anonymous users are granted permission to make calls to this endpoint.

* **URL**

  ``/api/v1/public/organizations``


* **Method**

  ``GET``


* **URL Params**

   * page


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

  ```json
  [
      {
          "description": "The company",
          "id": 1,
          "name": "Mika Software Corporation"
      }
  ]
  ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/public/organizations page==1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Get Profile
The API endpoint used to get the *user profile details*. Only the *profile* of the
  *authenticated user* is returned.

* **URL**

  ``/api/v1/profile``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "user_id": 1
    }
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/profile Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Create Organization
Creates an *organization* by an *authenticated user* in our system. Please note the following rules:

* *User* is able to create only a single *organization* and cannot create anymore.
* *User* cannot be an employee of an *organization*.

Once an *organization* has been created, the staff of [**Lucha Comics** ](https://luchacomics.com/) must approve the *organization* before it becomes publicly viewable on Comics Cantina.

* **URL**

  ``/api/v1/organizations``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

    * name
    * description
    * email
    * street_address
    * street_address_extra
    * city
    * province
    * country
    * currency - optional
    * language - optional
    * website - optional
    * phone - optional
    * fax - optional


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "city": "London",
        "country": "Canada",
        "description": "The company",
        "email": "bart@mikasoftware.com",
        "id": 1,
        "name": "Mika Software Corporation",
        "owner_id": 1,
        "province": "Ontario",
        "status": 1,
        "street_address": "111-204 Infinite Loop Road"
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Name is not unique.",
        "status": "Invalid request."
    }
    ```

  OR

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Cannot create organization because you have already created an organization. You are allowed to only have one organization per account.",
        "status": "Invalid request."
    }
    ```

  OR

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "Cannot create organization because you are an employee. Please create a new account if you want to create an organization.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/organizations \
    Authorization:"Bearer $COMICS_WS_API_TOKEN" \
    name="Mika Software Corporation" \
    description="The company" \
    email="bart@mikasoftware.com" \
    street_address="111-204 Infinite Loop Road" \
    city="London" province="Ontario" country="Canada"
  ```


## List Organizations
Returns paginated list of all the *organizations* if the *authenticated user* is a staff member of [**Lucha Comics** ](https://luchacomics.com/).

* **URL**

  ``/api/v1/organizations``


* **Method**

  ``GET``


* **URL Params**

  * page


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    [
        {
            "description": "The company",
            "id": 1,
            "name": "Mika Software Corporation"
        }
    ]
    ```


* **Error Response**

  * **Code:** 401
  * **Content:**

    ```json
    You are not a staff member.
    ```


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/organizations page==1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Retrieve Organization
Returns the *organization* details. Only *authenticated users* which meet the following criteria are allowed to access this endpoint:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then they are automatically granted access.

* **URL**

  ``/api/v1/organization/<organization_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "city": "London",
        "country": "Canada",
        "description": "The company",
        "email": "bart@mikasoftware.com",
        "id": 1,
        "name": "Mika Software",
        "owner_id": 1,
        "province": "Ontario",
        "status": 1,
        "street_address": "111-204 Infinite Loop Road"
    }
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/organization/1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Update Organization

**TODO: IMPLEMENT**


## Create Store
The API endpoint used to create a *Store* in Comics Cantina by an *authenticated user* belongoing to a specific *Organization*.

* **URL**

  ``/api/v1/stores``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * organization_id
  * name
  * description
  * email
  * street_address
  * street_address_extra
  * city
  * province
  * country
  * currency - optional
  * language - optional
  * website - optional
  * phone - optional
  * fax - optional


* **Success Response**

  * **Code:** 200
  * **Content:**

  ```json
  {
      "city": "London",
      "country": "Canada",
      "description": "The brick and morter comics store.",
      "email": "bart@mikasoftware.com",
      "id": 1,
      "name": "Main Store",
      "organization_id": 1,
      "province": "Ontario",
      "status": 1,
      "street_address": "111-204 Infinite Loop Road"
  }
  ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
      "error": "`organization_id` is invalid - either does not exist or you are not a member of it.",
      "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/stores \
    Authorization:"Bearer $COMICS_WS_API_TOKEN" \
    name="Downtown Main Store" \
    description="The brick and morter comics store in downtown." \
    email="bart@mikasoftware.com" \
    street_address="111-204 Infinite Loop Road" \
    city="London" \
    province="Ontario" \
    country="Canada" \
    organization_id=1
  ```


## List Stores
Returns paginated list of all the *stores* that meet any of these criteria for the *authenticated user* that made the call:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then all *organizations* get listed regardless of membership or status.

* **URL**

  ``/api/v1/stores``


* **Method**

  ``GET``


* **URL Params**

  * page


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    [
        {
            "description": "The company",
            "id": 1,
            "name": "Mika Software Corporation"
        }
    ]
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/stores page==1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Retrieve Store
Returns the *store* details. Only *authenticated users* which meet the following criteria are allowed to access this endpoint:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then they are automatically granted access.

* **URL**

  ``/api/v1/store/<store_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    [
        {
            "description": "The brick and morter comics store in downtown.",
            "id": 1,
            "name": "Downtown Main Store",
            "organization_id": 1
        }
    ]
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/store/1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Update Store

**TODO: IMPLEMENT**


## Create Category
The API endpoint used to create a *Category* in Comics Cantina by an *authenticated user* belongoing to a specific *Store*.

* **URL**

  ``/api/v1/categories``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

  ```json
  TODO: IMPLEMENT
  ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "`organization_id` is invalid - either does not exist or you are not a member of it.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/categories \
    Authorization:"Bearer $COMICS_WS_API_TOKEN" \
    name="Sci-Fi" \
    short_description="Some short description." \
    long_description="Long description." \
    organization_id=1
  ```


## List Categories
Returns paginated list of all the *categories* that meet any of these criteria for the *authenticated user* that made the call:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then all *organizations* get listed regardless of membership or status.

* **URL**

  ``/api/v1/categories``


* **Method**

  ``GET``


* **URL Params**

  * page


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    [
        {
            "id": 2,
            "long_description": "Long description.",
            "name": "Sci-Fi",
            "short_description": "Some short description."
        }
    ]
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/categories Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Retrieve Category
Returns the *category* details. Only *authenticated users* which meet the following criteria are allowed to access this endpoint:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then they are automatically granted access.

* **URL**

  ``/api/v1/category/<category_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "id": 1,
        "long_description": "Long description.",
        "name": "Sci-Fi",
        "organization_id": 1,
        "short_description": "Some short description."
    }
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/category/1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Update Category

**TODO: IMPLEMENT**


## Create Supplier
The API endpoint used to create a *Supplier* in Comics Cantina by an *authenticated user* belongoing to a specific *Store*.

* **URL**

  ``/api/v1/suppliers``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

  ```json
  {
      "id": 1,
      "name": "Comic Book Suppliers LTD.",
      "organization_id": 1
  }
  ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
        "error": "`organization_id` is invalid - either does not exist or you are not a member of it.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/suppliers \
    Authorization:"Bearer $COMICS_WS_API_TOKEN" \
    name="Comic Book Suppliers LTD." \
    organization_id=1
  ```


## List Categories
Returns paginated list of all the *suppliers* that meet any of these criteria for the *authenticated user* that made the call:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then all *organizations* get listed regardless of membership or status.

* **URL**

  ``/api/v1/suppliers``


* **Method**

  ``GET``


* **URL Params**

  * page


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    [
        {
            "id": 2,
            "name": "Comic Book Suppliers LTD.",
            "organization_id": 1
        }
    ]
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/suppliers Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Retrieve Supplier
Returns the *category* details. Only *authenticated users* which meet the following criteria are allowed to access this endpoint:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then they are automatically granted access.

* **URL**

  ``/api/v1/category/<category_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    {
        "id": 1,
        "name": "Comic Book Suppliers LTD.",
        "organization_id": 1
    }
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/supplier/1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Update Supplier

**TODO: IMPLEMENT**


## Create Product
The API endpoint used to create a *Product* in Comics Cantina by an *authenticated user* belongoing to a specific *Store*.

* **URL**

  ``/api/v1/products``


* **Method**

  ``POST``


* **URL Params**

  None


* **Data Params**

  * name
  * organization_id
  * store_id


* **Success Response**

  * **Code:** 200
  * **Content:**

  ```json
  TODO: IMPLEMENT
  ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```json
    {
      "error": "`organization_id` is invalid - either does not exist or you are not a member of it.",
      "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```bash
  $ http post 127.0.0.1:8080/api/v1/products \
    Authorization:"Bearer $COMICS_WS_API_TOKEN" \
    name="Winterworld" \
    organization_id=1 \
    store_id=1
  ```


## List Products
Returns paginated list of all the *products* that meet any of these criteria for the *authenticated user* that made the call:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then all *organizations* get listed regardless of membership or status.

* **URL**

  ``/api/v1/products``


* **Method**

  ``GET``


* **URL Params**

  * page
  * store_id


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    TODO: IMPLEMENT
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/products page==1 Authorization:"Bearer $COMICS_WS_API_TOKEN"

  OR

  $ http get 127.0.0.1:8080/api/v1/products page==1 store_id==1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Retrieve Product
Returns the *product* details. Only *authenticated users* which meet the following criteria are allowed to access this endpoint:

  * *user* is the owner of the *organization*
  * *user* is an employee of the *organization*

It is important to note that if the *authenticated user* is staff member of [**Lucha Comics** ](https://luchacomics.com/) then they are automatically granted access.

* **URL**

  ``/api/v1/product/<product_id>``


* **Method**

  ``GET``


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```json
    TODO: IMPLEMENT
    ```


* **Error Response**

  * None


* **Sample Call**

  ```bash
  $ http get 127.0.0.1:8080/api/v1/product/1 Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```


## Update Product

**TODO: IMPLEMENT**
