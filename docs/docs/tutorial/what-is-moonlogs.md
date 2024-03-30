# What is Moonlogs

Moonlogs is a specialized logging solution designed with a primary focus on capturing business-event logs. It excels at providing insights into business processes, user interactions, and custom events tailored for your applications. It even comes with a built-in web interface for convenient searching of the necessary events.

## Motivation

Businesses generate vast amounts of data from various events such as user interactions, transactions, integrations with third-party services, complicated business operations over time and so much more.

Unfortunately, as the number of features grows and the number of employees in the company increases, keeping track of all processes becomes increasingly difficult. This is especially true for non-trivial business operations that extend over time.

To address these challenges, companies address this using several approaches:
* Maintaining the knowledge base;
* Creating own logging solutions.


## What is wrong with maintaining knowledge base

While this approach is incredibly useful in itself, significantly aiding in onboarding new employees and helping to troubleshoot issues, it still has several drawbacks:

* **Navigating the knowledge base**

  In a vast amount of information, it can be quite problematic to find the necessary material.
  And sometimes it's not entirely clear what exactly needs to be searched for.

* **Knowledge can become outdated**

  With active development/support, processes within the application can change frequently: both in detail and in a global sense. Because of this, we may face two unpleasant moments: either the knowledge base needs to be changed instantly, consuming the time of developers and other employees, or we will have outdated data.
  Changes to the knowledge base are necessary in any case, but we are still not immune to discrepancies.

Therefore, maintaining a knowledge base, while being an excellent practice and methodology, still does not completely solve the task.

## What is wrong with developing own logging solution

Developing your own logging solutions instead of using existing ones can have several potential drawbacks:

* **Time and effort**

  Creating a logging solution from scratch requires time, effort, and resources that could be allocated to other aspects of your project.

* **Complexity and Maintenance**

  Building and maintaining a custom logging solution can introduce some complexity. You'll need to ensure that your logging solution is robust, scalable, and capable of handling multiple scenarios and even edge cases. Also, the burden of updates, bug fixes, and the development of additional features falls entirely on your shoulders.

* **Scalability**

  As the project grows, or even as new projects are added to the business, the logging requirements may also evolve. In this regard, you'll need to take an effor to support large-scale deployments and distributed architectures, spending time on this instead of solving the business tasks.

Custom logging solutions, while a fairly common approach, require a lot of resources and time from company employees, yet still do not fully address the task.

## And what task we need to solve?

The goal — is to bring transparency.

This means that at any given moment, employees from various departments should have the ability to promptly access a sufficient amount of clear information that fully describes a specific business process.

Moreover, the method of obtaining such information should be as straightforward as possible, and the information itself should be properly structured and easy to understand, even for non-technical employees.

## And so we developed Moonlogs

We developed Moonlogs to address the above problems by providing a centralized tool for storing, aggregating, and presenting information that reflects the details of your business processes.

By offering a user-friendly interface, structured data organization, and real-time insights, Moonlogs enables companies to enhance visibility, streamline knowledge management, adapt to dynamic business environments, facilitate employee onboarding and training, and make informed decisions based on up-to-date information.

However, it's important to note that Moonlogs **is not a replacement** for a knowledge base, system logging, or application performance monitoring (APM). Instead, it complements these existing tools and practices as part of a comprehensive approach to making your application transparent and predictable.

## Main features

As the result, Moonlogs offers you a set of features to fulfill the desired goal:

### Meta-groups of events, based on domain areas.

You can create separate meta-groups (schemas) to categorize events by domain areas. For instance, create schemas for the checkout process, user access setting changes, and Uber Eats integration. Events within each schema are recorded independently, facilitating efficient event retrieval.

### Query-based log subgrouping

Group events within a schema based on specified queries to enhance information integrity. This not only simplifies searchability but also ensures unrelated events remain separate even if in the same schema.

### Convenient schema-based filtering

Moonlogs generates convenient filters on the web interface for each schema, simplifying event search by allowing users to simply input values. This eliminates the complexity of composing queries with an undefined set of parameters, making it user-friendly, especially for non-technical personnel.

### Flexible event retention time

Specify varying retention times for each schema to align with specific business needs. For instance, set a 7-day retention time for event in the "Glovo integration" schema, while events in the "User's rights change history" schema can be stored indefinitely. Adjust these settings dynamically as business requirements evolve.

### Granular access control with tags

Create and assign tags to schemas and users, enabling granular access control. Define access privileges based on tags, ensuring that users can only access the schemas and events relevant to their responsibilities. This feature provides an additional layer of security and customization in managing access to events data.

## What's next?

* Go to [Installation](/tutorial/install) section to find out how to install the Moonlogs on your system

* Check out [Introduction to the Web UI](/web-ui/introduction) section to familiarize yourself with the Moonlogs built-in web-interface
